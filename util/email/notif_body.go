package email

import "fmt"

type NotifBody string

const (
	NotifBodyLogin  NotifBody = "Your account was just accessed from a %v. If you did not initiate this login, please reset your password immediately."
	NotifBodySignUp NotifBody = `<html>
	<body>
	  <p>Hello,</p>
	  <p>Thank you for signing up for our service! In order to complete your registration, please click the link below to verify your email address:</p>
	  <p>Verify your email address</p>
	  <p>Token: <b>%s</b> </p>
	  <p><a href="%s/user/signup/%s">Verify your email address</a></p>
	  <p>If you did not request this verification email, please ignore this message.</p>
	  <p>Best regards,<br>Project Team</p>
	</body>
  </html>`

	NotifBodyChangePass NotifBody = `<html>
  <body>
    <p>Hi there,</p>
    <p>We wanted to let you know that your password was recently changed from the IP address %s.</p>
    <p>If you did not request this change, please contact us immediately to secure your account.</p>
	<a href="%s/user/change-password/%s">Change Password</a>
    <p>Best regards,</p>
    <p> <br>Project Team</br> </p>
  </body>
</html>
`
	NotifBodyPassChanging NotifBody = `<body>
    <h1>Password Change Notification</h1>
    <p>Hello,</p>
    <p>This is a notification to let you know that your password has been successfully changed.</p>
    <p>If you did not initiate this change, please contact support immediately.</p>
    <p>Thanks,</p>
    <p>The Support Team</p>
</body>`
)

// here enum for the result email body
// for signup, the format is [0] = token, [1] = base url [2] = uuid
func (n NotifBody) Format(Format ...string) string {
	switch n {
	case NotifBodyLogin:
		return fmt.Sprintf(string(n), Format[0])
	case NotifBodySignUp:
		return fmt.Sprintf(string(n), Format[0], Format[1], Format[2])
	case NotifBodyChangePass:
		return fmt.Sprintf(string(n), Format[0], Format[1], Format[2])
	default:
		return string(n)
	}
}
func (n NotifBody) String() string {
	return string(n)
}
