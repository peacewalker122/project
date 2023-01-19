package user

type CreateUserParam struct {
	Username       string `json:"username" form:"username" validate:"required,min=4,max=100"`
	HashedPassword string `json:"password" form:"password" validate:"required,min=6,max=100"`
	FullName       string `json:"fullname" form:"fullname" validate:"required,min=3,max=100"`
	Email          string `json:"email" form:"email" validate:"required,email"`
}

type CreatingUser struct {
	Token int `json:"token" form:"token" query:"token" validate:"required"`
}

type LoginParams struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type ValidateChangePassRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" form:"password" validate:"required"`
}
