package email

type NotifHeader string

const (
	NotifHeaderLogin      NotifHeader = "Login Notification"
	NotifHeaderSignUp     NotifHeader = "Welcome to Project!"
	NotifHeaderChangePass NotifHeader = "Password Changed"
)
