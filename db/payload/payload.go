package payload

import (
	"database/sql"

	"github.com/peacewalker122/project/db/payload/model"
)

func NewPayload(sql *sql.DB) Payload {
	return model.NewModel(sql)
}

type Payload interface {
	model.Model
}
