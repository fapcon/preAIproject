package models

import "time"

type UserBaned struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Reason        string    `json:"reason"`
	BanStartedAt  time.Time `json:"ban_started_at"`
	BanFinishedAt time.Time `json:"ban_finished_at"`
}
