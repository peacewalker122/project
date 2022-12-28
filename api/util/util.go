package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	notifquery "github.com/peacewalker122/project/db/model/notif_query"
	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
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
}

func NewApiUtil(store db.Store, redis redis.Store, cfg util.Config) UtilTools {
	return &utilTools{
		store: store,
		redis: redis,
		cfg:   cfg,
	}
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

func (s *utilTools) SendEmail(params ...string) error {
	str := string(NotifBodyLogin.Format(params[1]))

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.cfg.Email)
	mailer.SetHeader("To", params[0])
	mailer.SetAddressHeader("Cc", params[0], "Admin")
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

	err = s.SendEmail(params.Params[0], params.Params[1])
	if err != nil {
		return err
	}
	return nil
}
