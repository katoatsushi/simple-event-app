package models

import "time"

type Event struct {
	ID          uint64        `json:"id"`
	EventName   string        `json:"event_name"`
	StartTime   time.Time     `json:"start_time"`
	Cost        int           `json:"cost"`
	Detail      string        `gorm:"type:text;" json:"detail"`
	EventSubTag []EventSubTag `json:"event_sub_tag"`
	TagID       uint64        `json:"tag_id"`
	Tag         Tag           `json:"tag"`
}
