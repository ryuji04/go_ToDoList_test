package server

import (
	"net/http"
	"test/controller"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	gin.SetMode(gin.ReleaseMode)
	controller := controller.NewTask()
	tasks := controller.GetAll()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "title",
		"tasks": tasks,
	})
}
