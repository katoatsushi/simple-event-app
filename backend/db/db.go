package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// database’s driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/simple-event-app/backend/config"
	"github.com/simple-event-app/backend/models"
)

var Tables = []interface{}{
	&models.Event{},
	&models.EventSubTag{},
	&models.SubTag{},
	&models.Tag{},
}

func Init() *gorm.DB {

	url := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetConfig().Get("db_user"),
		config.GetConfig().Get("db_password"),
		config.GetConfig().Get("db_host"),
		config.GetConfig().Get("db_name"))
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(url)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")

	db, err := gorm.Open("mysql", url)

	if err != nil {
		panic("failed to connect database")
	}

	if config.GetConfig().Get("env") == "development" {
		db.LogMode(true)

		db.AutoMigrate(Tables...)
		AddForeignKey(db)

	}
	return db
}

func AddForeignKey(db *gorm.DB) {
	db.Model(&models.EventSubTag{}).AddForeignKey("event_id", "events(id)", "CASCADE", "CASCADE")
	db.Model(&models.EventSubTag{}).AddForeignKey("sub_tag_id", "sub_tags(id)", "CASCADE", "CASCADE")
	db.Model(&models.Event{}).AddForeignKey("tag_id", "tags(id)", "CASCADE", "CASCADE")
}
