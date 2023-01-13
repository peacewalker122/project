package notifquery

import "time"

type NotifParams struct {
	AccountID    []int64
	Username     string
	NotifType    string
	NotifTitle   string
	NotifContent string
	NotifTime    time.Time
}
