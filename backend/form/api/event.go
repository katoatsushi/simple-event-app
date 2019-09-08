package api

import "time"

type EventCreateIn struct {
	EventName      string    `json"event_name"`
	StartTime      time.Time `json"start_time"`
	Cost           int       `json"cost"`
	Detail         string    `gorm:"type:text;" json"detail"`
	EventSubTagIDs []uint64  `json"event_sub_tag_ids"`
	TagID          uint64    `json"tag_id"`
}
