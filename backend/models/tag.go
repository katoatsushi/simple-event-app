package models

type Tag struct {
	ID uint64 `json"id"`
	//Name   constants.TagName　`json"name"`
	Name   string `json"name"`
	Detail string `json"detail"`
}
