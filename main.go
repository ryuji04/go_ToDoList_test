package main

import (
	"fmt"
	"net/http"

	"strconv"
	"test/controller"
	"test/server"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("template/*html")
	router.GET("/", server.GetAll)

	router.POST("/", func(c *gin.Context) {
		text := c.PostForm("text")
		ctrl := controller.NewTask()
		ctrl.Create(text)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.GET("/:id", func(c *gin.Context) {
		var id, _ = strconv.Atoi(c.Param("id"))
		ctrl := controller.NewTask()
		tasks := ctrl.GetById(id)
		fmt.Println("tasks:", tasks)
		c.HTML(http.StatusOK, "index.html", gin.H{"tasks": tasks})
	})
	router.Run(":8080")
}
