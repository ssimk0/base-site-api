package models

import "time"

type Announcement struct {
	DatabaseModel
	Message  string     `json:"message,omitempty"`
	ExpireAt *time.Time `json:"expire_at,omitempty"`
}

type AnnouncementRequest struct {
	Message  string     `json:"message"`
	ExpireAt *time.Time `json:"expire_at"`
}
