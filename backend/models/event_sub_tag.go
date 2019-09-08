package models

type EventSubTag struct {
	ID       uint64 `json"id"`
	EventID  uint64 `json"event_id"`
	Event    Event  `json"-"`
	SubTagID uint64 `json"sub_tag_id"`
	SubTag   SubTag `json"-"`
}
