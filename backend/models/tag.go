package models

type Tag struct {
	ID uint64 `json:"id"`
	//Name   constants.TagName　`json"name"`
	Name     string `json:"name"`
	TagIntro string `gorm:"type:text;" json:"tag_intro"`
	Detail   string `gorm:"type:text;" json:"detail"`
}
