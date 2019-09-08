package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simple-event-app/backend/form/api"
	"github.com/simple-event-app/backend/models"
	"github.com/simple-event-app/backend/module"
)

func CreateTag(router *gin.RouterGroup) {
	router.POST("tag", func(c *gin.Context) {
		in := new(api.TagCreateIn)
		if err := c.Bind(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tag := &models.Tag{
			Name:     in.Name,
			TagIntro: in.TagIntro,
			Detail:   in.Detail,
		}
		err := module.Configure.Repository.TagRepository.Create(tag)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error occured"})
			return
		}
		c.JSON(http.StatusOK, "success")
	})
}
