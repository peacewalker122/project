package contract

import (
	"github.com/peacewalker122/project/contract/auth"
	"github.com/peacewalker122/project/contract/user"
)

type Contract interface {
	user.UserContract
	auth.AuthContract
}
