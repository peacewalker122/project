package notifquery

import (
	"context"

	"github.com/google/uuid"
	"github.com/peacewalker122/project/db/repository/postgres/ent"
	notif "github.com/peacewalker122/project/db/repository/postgres/ent/accountnotifs"
)

// NotifQuery returns a query builder for Notif.
type NotifQuery interface {
	CreateNotif(ctx context.Context, Params *NotifParams) (*ent.AccountNotifs, error)
	GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent.AccountNotifs, error)
	CreateNotifUsername(ctx context.Context, Params *NotifParams) (*ent.AccountNotifs, error)
	DeleteNotif(ctx context.Context, notifID uuid.UUID) error
}

type NotifsQueries struct {
	*ent.Client
}

// CreateNotif implements NotifsQueries
func (n *NotifsQueries) CreateNotif(ctx context.Context, Params *NotifParams) (*ent.AccountNotifs, error) {
	var (
		err error
		res *ent.AccountNotifs
	)
	uid := uuid.New() // we create this for make sure if we have a plenty notif it will called at the same time then sending it
	for _, v := range Params.AccountID {
		res, err = n.Client.AccountNotifs.
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

func (n *NotifsQueries) CreateNotifUsername(ctx context.Context, Params *NotifParams) (*ent.AccountNotifs, error) {
	uid := uuid.New() // we create this for make sure if we have a plenty notif it will called at the same time then sending it
	res, err := n.Client.AccountNotifs.
		Create().
		SetID(uid).
		SetUsername(Params.Username).
		SetNotifType(Params.NotifType).
		SetNotifTime(Params.NotifTime).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetNotifByAccount implements NotifsQueries
func (n *NotifsQueries) GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent.AccountNotifs, error) {
	res, err := n.Client.AccountNotifs.
		Query().
		Where(notif.AccountID(accountID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteNotif implements NotifsQueries
func (n *NotifsQueries) DeleteNotif(ctx context.Context, notifID uuid.UUID) error {
	_, err := n.Client.AccountNotifs.
		Delete().
		Where(notif.ID(notifID)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewNotifQuery(Driver *ent.Client) *NotifsQueries {
	return &NotifsQueries{
		Client: Driver,
	}
}
