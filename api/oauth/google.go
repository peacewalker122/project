package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/db/model/tokens"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	domain   = "https://58b0-2001-448a-6021-850e-2d31-3650-e295-8bff.ap.ngrok.io"
	redirect = "/oauth/google/callback"
)

var (
	GoogleCfg = oauth2.Config{}
)

func (s *Handler) GoogleVerif(c echo.Context) error {
	GoogleCfg = oauth2.Config{
		ClientID:     s.config.GoogleClientID,
		ClientSecret: s.config.GoogleClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", domain, redirect),
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
	errchan := make(chan error, 2)

	wg.Add(1)
	go func(url string) {
		defer wg.Done()
		resp, err := client.Get(url)
		if err != nil {
			errchan <- err
		}

		err = json.NewDecoder(resp.Body).Decode(&payload)
		if err != nil {
			errchan <- err
		}
	}(url)

	if len(errchan) > 0 {
		for v := range errchan {
			return c.JSON(http.StatusBadRequest, v.Error())
		}
	}

	if ok, err := s.store.IsTokenExist(ctx, payload.Email); !ok && err == nil {
		_, err = s.store.SetToken(ctx, &tokens.TokensParams{
			Email:        payload.Email,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			TokenType:    token.TokenType,
			ExpiresIn:    token.Expiry,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	wg.Wait()
	return c.JSON(http.StatusOK, payload)
}
