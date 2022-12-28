package api

import "fmt"

type NotifType string

const (
	NotifTypeLike    NotifType = "like"
	NotifTypeComment NotifType = "comment"
	NotifTypeFollow  NotifType = "follow"
	NotifTypeMention NotifType = "mention"
	NotifTypeShare   NotifType = "share"
	NotifTypeRetweet NotifType = "retweet"
	NotifTypeLogin   NotifType = "login"
	NotifTypeAds     NotifType = "ads"
)

type NotifBody string

const (
	NotifBodyLogin  NotifBody = "Your account was just accessed from a %v. If you did not initiate this login, please reset your password immediately."
	NotifBodySignUp NotifBody = `Hello!

	Thank you for signing up for our service. We're excited to have you on board!
	
	To complete your sign up, please click the link below to verify your email address:
	
	%v
	
	Once you've verified your email, you can start using our service right away. Here are just a few of the features you can look forward to:
	
	Get started now!
	
	Best regards,
	Project Team`
)

// here enum for the response email body
func (n NotifBody) Format(Format string) string {
	switch n {
	case NotifBodyLogin:
		return fmt.Sprintf(string(n), Format)
	case NotifBodySignUp:
		return fmt.Sprintf(string(n), Format)
	default:
		return string(n)
	}
}
func (n NotifBody) String() string {
	return string(n)
}

type NotifHeader string

const (
	NotifHeaderLogin  NotifHeader = "Login Notification"
	NotifHeaderSignUp NotifHeader = "Welcome to Project!"
)
