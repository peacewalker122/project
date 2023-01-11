package payload

import (
	"database/sql"

	"github.com/peacewalker122/project/db/payload/model"
)

func NewPayload(sql ...*sql.DB) Payload {
	return model.NewModel(sql[0], sql[1])
}

type Payload interface {
	model.Model
}
