package email

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	notifquery "github.com/peacewalker122/project/db/repository/postgres/payload/model/notif_query"
	"github.com/peacewalker122/project/util"
	"gopkg.in/gomail.v2"
)

type SendEmail struct {
	AccountID []int64
	Params    []string
	Type      NotifType
	TimeSend  time.Time
}

// NotifType is the type of notification
// params[0] = email address, params[1] = ip address/token, params[2] = uuid
// params[2] is required for signup only
func (s *EmailHelper) SendEmail(types NotifType, params ...string) error {
	var str string
	switch types {
	case NotifTypeLogin:
		str = string(NotifBodyLogin.Format(params[1]))
	case NotifTypeSignUp:
		str = string(NotifBodySignUp.Format(params[1], params[2]))
	case NotifTypeChangePass:
		str = string(NotifBodyChangePass.Format(params[1], params[2]))
	case NotifTypePassChanging:
		str = string(NotifBodyPassChanging.Format(params[1]))
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.cfg.Email)
	mailer.SetHeader("To", params[0])
	mailer.SetAddressHeader("Cc", params[0], "Login")
	mailer.SetHeader("Subject", string(NotifHeaderLogin))
	mailer.SetBody("text/html", fmt.Sprintf("<p>%s</p>", str))

	//log.Panic(fmt.Sprintf("email: %s, pass: %s", s.cfg.Email, s.cfg.EmailPass))

	dialer := gomail.NewDialer(s.cfg.EmailSMTP, 587, s.cfg.Email, s.cfg.EmailPass)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

func (s *EmailHelper) SendEmailWithNotif(ctx context.Context, params SendEmail) error {

	_, err := s.postgre.CreateNotif(ctx, &notifquery.NotifParams{
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

func (s *EmailHelper) CreateEmailAuth(ctx context.Context, email string) (uuid.UUID, error) {
	// uid indicate for the email auth session
	uid := uuid.New()

	// to make token consist of 6 digits
	token := util.Randomint(100000, 999999)

	err := s.redis.SetOne(ctx, uid.String(), token, 3*time.Minute)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to set token to redis: %w", err)
	}

	err = s.SendEmail(NotifTypeSignUp, email, fmt.Sprintf("%d", token), uid.String())
	if err != nil {
		return uuid.UUID{}, err
	}

	return uid, nil
}

func (s *EmailHelper) ChangePasswordAuth(ctx context.Context, params SendEmail) error {
	// uid indicate for the email auth session
	//var err error

	acc, err := s.postgre.GetAccountByEmail(ctx, params.Params[0])
	if err != nil {
		return err
	}

	_, err = s.postgre.CreateNotifUsername(ctx, &notifquery.NotifParams{
		AccountID: []int64{acc.ID},
		NotifType: string(params.Type),
		NotifTime: params.TimeSend,
	})
	if err != nil {
		return err
	}

	err = s.SendEmail(NotifTypeChangePass, params.Params[0], params.Params[1], params.Params[2])
	if err != nil {
		return err
	}

	return nil
}

func (s *EmailHelper) VerifyEmailAuth(ctx context.Context, uid string, token int) (bool, error) {
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

	// delete token from redis
	s.redis.Del(ctx, uid)

	return true, err
}
