package email

type NotifType string

const (
	NotifTypeLike         NotifType = "like"
	NotifTypeComment      NotifType = "comment"
	NotifTypeFollow       NotifType = "follow"
	NotifTypeMention      NotifType = "mention"
	NotifTypeShare        NotifType = "share"
	NotifTypeRetweet      NotifType = "retweet"
	NotifTypeLogin        NotifType = "login"
	NotifTypeAds          NotifType = "ads"
	NotifTypeSignUp       NotifType = "signup"
	NotifTypeChangePass   NotifType = "password change"
	NotifTypePassChanging NotifType = "password-changing"
)
