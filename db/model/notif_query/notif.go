package notifquery

import (
	"context"

	"github.com/google/uuid"
	"github.com/peacewalker122/project/db/ent"
	notif "github.com/peacewalker122/project/db/ent/accountnotif"
)

// NotifQuery returns a query builder for Notif.
type NotifQuery interface {
	CreateNotif(ctx context.Context, Params *NotifParams) (*ent.AccountNotif, error)
	GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent.AccountNotif, error)
	DeleteNotif(ctx context.Context, notifID uuid.UUID) error
}

type NotifsQueries struct {
	*ent.Client
}

// CreateNotif implements NotifsQueries
func (n *NotifsQueries) CreateNotif(ctx context.Context, Params *NotifParams) (*ent.AccountNotif, error) {
	var (
		err error
		res *ent.AccountNotif
	)
	uid := uuid.New() // we create this for make sure if we have a plenty notif it will called at the same time then sending it
	for _, v := range Params.AccountID {
		res, err = n.Client.AccountNotif.
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

// GetNotifByAccount implements NotifsQueries
func (n *NotifsQueries) GetNotifByAccount(ctx context.Context, accountID int64) ([]*ent.AccountNotif, error) {
	res, err := n.Client.AccountNotif.
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
	_, err := n.Client.AccountNotif.
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
