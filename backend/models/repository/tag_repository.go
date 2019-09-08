package repository

import "github.com/simple-event-app/backend/models"

type TagRepository interface {
	Create(tag *models.Tag) error
}
