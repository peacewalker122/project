package account

import "database/sql"

type AccountParam struct {
	Owner     string         `json:"owner"`
	IsPrivate bool           `json:"is_private"`
	PhotoDir  sql.NullString `json:"photo_dir"`
}
