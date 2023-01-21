package sqlcTX

import (
	"database/sql"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/TX/account"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/TX/post"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/TX/user"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

type SQLCTX interface {
	post.PostDBTX
	user.UserDBTX
	account.AccountDBTX
}

type Tx struct {
	Store *db.SQLStore
	*post.PostTx
	*user.UserTx
	*account.AccountTx
}

func NewTx(project *sql.DB) *Tx {
	res := &Tx{
		Store: db.NewStore(project),
	}
	res.PostTx = post.NewPostTx(res.Store)
	res.UserTx = user.NewUserTx(res.Store)
	res.AccountTx = account.NewAccountTx(res.Store)
	return res
}
