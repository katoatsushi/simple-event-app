package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/simple-event-app/backend/models"
)

type TagMysql struct {
	DbCon *gorm.DB
}

func (tagMysql *TagMysql) Create(tag *models.Tag) error {
	if err := tagMysql.DbCon.Create(tag); err.Error != nil {
		return err.Error
	}
	return nil
}
