package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/simple-event-app/backend/controller"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	mainRouter := router.Group("evet_app")
	controller.CreateEvent(mainRouter)

	return router
}
