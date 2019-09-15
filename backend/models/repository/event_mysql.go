package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/simple-event-app/backend/models"
)

type EventMysql struct {
	DbCon *gorm.DB
}

func (eventMysql *EventMysql) Create(event *models.Event) error {

	tx := eventMysql.DbCon.Begin()

	if err := tx.Create(event); err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	tx.Commit()
	return nil
}
