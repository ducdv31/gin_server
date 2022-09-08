package main

import (
	"gin_server/controller"
	"gin_server/middleware"
	"gin_server/service"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.New()
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
	)

	/**
	 * Router
	 */
	server.POST("/video", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
