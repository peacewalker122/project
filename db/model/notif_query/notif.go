package notifquery

import (
	"context"
	"database/sql"
	"log"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/db/ent"
	"github.com/peacewalker122/project/db/ent/notif"
)

// NotifQuery returns a query builder for Notif.
type NotifQuery interface {
	CreateNotif(ctx context.Context, Params *NotifParams) (*ent.Notif, error)
	GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent.Notif, error)
	DeleteNotif(ctx context.Context, notifID uuid.UUID) error
}

type notifQuery struct {
	*ent.Client
}

// CreateNotif implements NotifQuery
func (n *notifQuery) CreateNotif(ctx context.Context, Params *NotifParams) (*ent.Notif, error) {
	var (
		err error
		res *ent.Notif
	)
	uid := uuid.New() // we create this for make sure if we have a plenty notif it will called at the same time then sending it
	for _, v := range Params.AccountID {
		res, err = n.Client.Notif.
			Create().
			SetID(uid).
			SetAccountID(v).
			SetNotifType(Params.NotifType).
			SetNotifTime(Params.NotifTime).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetNotifByAccount implements NotifQuery
func (n *notifQuery) GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent.Notif, error) {
	res, err := n.Client.Notif.
		Query().
		Where(notif.AccountID(accountID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteNotif implements NotifQuery
func (n *notifQuery) DeleteNotif(ctx context.Context, notifID uuid.UUID) error {
	_, err := n.Client.Notif.
		Delete().
		Where(notif.ID(notifID)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewNotifQuery(db string) NotifQuery {
	sql, err := sql.Open("postgres", db)
	if err != nil {
		log.Panic(err.Error())
	}
	drv := entsql.OpenDB("postgres", sql)
	//defer sql.Close()

	return &notifQuery{
		Client: ent.NewClient(ent.Driver(drv)),
	}
}
