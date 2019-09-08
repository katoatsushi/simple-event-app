package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/simple-event-app/backend/controller/forms"
	"github.com/simple-event-app/backend/form/api"
	"github.com/simple-event-app/backend/module"
)

func CreateEvent(router *gin.RouterGroup) {
	router.POST("event", func(c *gin.Context) {
		in := new(api.EventCreateIn)
		if err := c.Bind(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		event := forms.NewEventForCreate(in)
		err := module.Configure.Repository.EventRepository.Create(event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error occured"})
			return
		}
		c.JSON(http.StatusOK, "success")
	})
}
