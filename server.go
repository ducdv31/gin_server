package main

import (
	"gin_server/controller"
	"gin_server/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	server := gin.Default()
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
