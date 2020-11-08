package models

import "time"

type Announcement struct {
	Model
	Message  string     `json:"message,omitempty"`
	ExpireAt *time.Time `json:"expire_at,omitempty"`
}
