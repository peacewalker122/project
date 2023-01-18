package payload

import (
	"database/sql"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model"
)

func NewPayload(sql *sql.DB) Payload {
	return model.NewModel(sql)
}

type Payload interface {
	model.Model
}
