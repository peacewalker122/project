package util

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	notifquery "github.com/peacewalker122/project/db/model/notif_query"
	"github.com/peacewalker122/project/db/model/tokens"
	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
	"golang.org/x/oauth2"
	"gopkg.in/gomail.v2"
)

type utilTools struct {
	store db.Store
	redis redis.Store
	cfg   util.Config
}

type UtilTools interface {
	accountFeature
	SendEmailWithNotif(ctx context.Context, params SendEmail) error // make sure params.params is in order: email, ipAdrress, type
	CreateEmailAuth(ctx context.Context, email string) (uuid.UUID, error)
	VerifyEmailAuth(ctx context.Context, uid string, token int) (bool, error)
	TokenHelper(ctx context.Context, token oauth2.TokenSource) (*oauth2.Token, error)
}

func NewApiUtil(store db.Store, redis redis.Store, cfg util.Config) UtilTools {
	return &utilTools{
		store: store,
		redis: redis,
		cfg:   cfg,
	}
}

func (s *utilTools) TokenHelper(ctx context.Context, token oauth2.TokenSource) (*oauth2.Token, error) {
	t, err := token.Token()
	if err != nil {
		return nil, err
	}
	if t.Valid() {
		return t, nil
	}
	t, err = s.RefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *utilTools) RefreshToken(ctx context.Context, token oauth2.TokenSource) (*oauth2.Token, error) {
	var res oauth2.Token

	newToken, err := oauth2.ReuseTokenSource(&res, token).Token()
	if err != nil {
		return nil, err
	}

	err = s.store.UpdateToken(ctx, &tokens.TokensParams{
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
		ExpiresIn:    newToken.Expiry,
		TokenType:    newToken.TokenType,
	})

	return newToken, nil
}

func ValidateFileType(input multipart.File) error {
	byte := make([]byte, 512)
	if _, err := input.Read(byte); err != nil {
		return err
	}
	file := http.DetectContentType(byte)

	s := log.Default()
	s.Print(file)

	if file == "image/jpg" || file == "image/jpeg" || file == "image/gif" || file == "image/png" || file == "image/webp" || file == "video/mp4" {
		return nil
	}
	return errors.New("invalid type! must be jpeg/jpg/gif/png/mp4")
}

type SendEmail struct {
	AccountID []int64
	Params    []string
	Type      NotifType
	TimeSend  time.Time
}

// NotifType is the type of notification
// params[0] = email address, params[1] = ip address/token, params[2] = uuid
// params[2] is required for signup only
func (s *utilTools) SendEmail(types NotifType, params ...string) error {
	var str string
	mailer := gomail.NewMessage()
	switch types {
	case NotifTypeLogin:
		str = string(NotifBodyLogin.Format(params[1]))
	case NotifTypeSignUp:
		str = string(NotifBodySignUp.Format(params[1], params[2]))
	}

	mailer.SetHeader("From", s.cfg.Email)
	mailer.SetHeader("To", params[0])
	mailer.SetAddressHeader("Cc", params[0], "Login")
	mailer.SetHeader("Subject", string(NotifHeaderLogin))
	mailer.SetBody("text/html", fmt.Sprintf("<p>%s</p>", str))
	dialer := gomail.NewDialer("smtp-mail.outlook.com", 587, s.cfg.Email, s.cfg.EmailPass)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}

func (s *utilTools) SendEmailWithNotif(ctx context.Context, params SendEmail) error {

	_, err := s.store.CreateNotif(ctx, &notifquery.NotifParams{
		AccountID: params.AccountID,
		NotifType: string(params.Type),
		NotifTime: params.TimeSend,
	})
	if err != nil {
		return err
	}

	err = s.SendEmail(params.Type, params.Params[0], params.Params[1])
	if err != nil {
		return err
	}
	return nil
}

func (s *utilTools) CreateEmailAuth(ctx context.Context, email string) (uuid.UUID, error) {
	// uid indicate for the email auth session
	uid := uuid.New()

	// to make token consist of 6 digits
	token := util.Randomint(100000, 999999)

	err := s.redis.SetOne(ctx, uid.String(), token, 3*time.Minute)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = s.SendEmail(NotifTypeSignUp, email, fmt.Sprintf("%d", token), uid.String())
	if err != nil {
		return uuid.UUID{}, err
	}

	return uid, nil
}

func (s *utilTools) VerifyEmailAuth(ctx context.Context, uid string, token int) (bool, error) {
	// get token from redis
	redisToken, err := s.redis.Get(ctx, uid)
	if err != nil {
		return false, err
	}

	// convert redisToken to int
	redisTokenInt, _ := strconv.Atoi(redisToken)

	if token != redisTokenInt {
		return false, errors.New("token is not valid, make sure you enter the correct token")
	}

	return true, err
}
