package notifquery

import (
	"context"
	ent2 "github.com/peacewalker122/project/service/db/repository/postgres/ent"
	notif "github.com/peacewalker122/project/service/db/repository/postgres/ent/accountnotifs"

	"github.com/google/uuid"
)

// NotifQuery returns a query builder for Notif.
type NotifQuery interface {
	CreateNotif(ctx context.Context, Params *NotifParams) (*ent2.AccountNotifs, error)
	GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent2.AccountNotifs, error)
	CreateNotifUsername(ctx context.Context, Params *NotifParams) (*ent2.AccountNotifs, error)
	DeleteNotif(ctx context.Context, notifID uuid.UUID) error
}

type NotifsQueries struct {
	*ent2.Client
}

// CreateNotif implements NotifsQueries
func (n *NotifsQueries) CreateNotif(ctx context.Context, Params *NotifParams) (*ent2.AccountNotifs, error) {
	var (
		err error
		res *ent2.AccountNotifs
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

func (n *NotifsQueries) CreateNotifUsername(ctx context.Context, Params *NotifParams) (*ent2.AccountNotifs, error) {
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
func (n *NotifsQueries) GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent2.AccountNotifs, error) {
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

func NewNotifQuery(Driver *ent2.Client) *NotifsQueries {
	return &NotifsQueries{
		Client: Driver,
	}
}
