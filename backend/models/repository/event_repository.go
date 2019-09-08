package repository

import "github.com/simple-event-app/backend/models"

type EventRepository interface {
	Create(event *models.Event) error
}
