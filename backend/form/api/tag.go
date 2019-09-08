package api

type TagCreateIn struct {
	Name     string `json"name"`
	TagIntro string `gorm:"type:text;" json"tag_intro"`
	Detail   string `gorm:"type:text;" json"detail"`
}
