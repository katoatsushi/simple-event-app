package task

import (
	"path/filepath"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	database "github.com/simple-event-app/backend/db"
	"github.com/simple-event-app/backend/models"
)

var Seeds = struct {
	Tags    []models.Tag
	SubTags []models.SubTag
}{}

func CreateSampleData(db *gorm.DB) {
	filepaths, _ := filepath.Glob(filepath.Join("/go/src/github.com/simple-event-app/backend/db", "data", "*.json"))
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}

	// truncateTable
	truncateTables(db, database.Tables...)

	// add foreign
	database.AddForeignKey(db)

	for _, product := range Seeds.Tags {
		db.Create(&product)
	}

	for _, notice := range Seeds.SubTags {
		db.Create(&notice)
	}
}

func truncateTables(db *gorm.DB, tables ...interface{}) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	for _, table := range tables {
		if err := db.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}

		db.AutoMigrate(table)
	}
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
}
