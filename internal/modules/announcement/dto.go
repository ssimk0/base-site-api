package announcement

import "time"

type Announcement struct {
	Message  string     `json:"message"`
	ExpireAt *time.Time `json:"expire_at"`
}
