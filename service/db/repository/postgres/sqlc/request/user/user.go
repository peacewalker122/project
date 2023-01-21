package request

type CreateUserParamsTx struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Email    string                 `json:"email"`
	FullName string                 `json:"full_name"`
	Payload  map[string]interface{} `json:"payload"`
}
