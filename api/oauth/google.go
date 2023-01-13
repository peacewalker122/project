package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/account"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/users"
	"github.com/peacewalker122/project/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	redirect = "/oauth/google/callback"
)

var (
	GoogleCfg = oauth2.Config{}
)

func (s *Handler) GoogleVerif(c echo.Context) error {
	GoogleCfg = oauth2.Config{
		ClientID:     s.config.GoogleClientID,
		ClientSecret: s.config.GoogleClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", s.config.BaseURL, redirect),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	url := GoogleCfg.AuthCodeURL("state", oauth2.AccessTypeOffline)

	return c.Redirect(http.StatusFound, url)
}

func (s *Handler) GoogleToken(c echo.Context) error {
	var payload GoogleUser
	ctx := c.Request().Context()
	code := c.FormValue("code")

	token, err := GoogleCfg.Exchange(ctx, code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err = s.util.TokenHelper(ctx, GoogleCfg.TokenSource(ctx, token))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid token")
	}

	if token.AccessToken == "" {
		return c.JSON(http.StatusBadRequest, "invalid token")
	}

	client := GoogleCfg.Client(ctx, token)

	if client == nil {
		return c.JSON(http.StatusBadRequest, "invalid client")
	}

	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken)

	var wg sync.WaitGroup
	errchan := make(chan error, 1)

	wg.Add(1)
	go func(url string) {
		defer wg.Done()
		resp, err := client.Get(url)
		if err != nil {
			errchan <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errchan <- fmt.Errorf("status code: %d", resp.StatusCode)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(&payload)
		if err != nil {
			errchan <- err
			return
		}
	}(url)

	select {
	case err := <-errchan:
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	default:
	}
	wg.Wait()

	ok, err := s.store.IsTokenExist(ctx, payload.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if !ok {
		_, err = s.store.SetUsersOauth(ctx, &model.CreateUsersOauthParam{
			User: &users.UsersParam{
				Email:          payload.Email,
				FullName:       payload.Name,
				Username:       payload.Name,
				HashedPassword: "",
			},
			OauthToken: &tokens.TokensParams{
				Email:        payload.Email,
				AccessToken:  token.AccessToken,
				RefreshToken: token.RefreshToken,
				TokenType:    token.TokenType,
				ExpiresIn:    token.Expiry,
			},
			Account: &account.AccountParam{
				Owner:     payload.Email,
				IsPrivate: false,
				PhotoDir:  util.InputSqlString(payload.Picture),
			},
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"token":   token.AccessToken,
	})
}
