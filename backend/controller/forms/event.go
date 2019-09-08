package forms

import (
	"github.com/simple-event-app/backend/form/api"
	"github.com/simple-event-app/backend/models"
)

func NewEventForCreate(in *api.EventCreateIn) *models.Event {
	var subTags []models.EventSubTag
	for _, subTagID := range in.EventSubTagIDs {
		subTag := models.SubTag{ID: subTagID}
		subTags = append(subTags, models.EventSubTag{
			SubTagID: subTagID,
			SubTag:   subTag,
		})
	}

	tag := models.Tag{ID: in.TagID}
	model := &models.Event{
		EventName:   in.EventName,
		StartTime:   in.StartTime,
		Cost:        in.Cost,
		Detail:      in.Detail,
		EventSubTag: subTags,
		TagID:       in.TagID,
		Tag:         tag,
	}
	return model
}
