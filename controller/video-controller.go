package controller

import (
	"gin_server/entity"
	"gin_server/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(context *gin.Context) entity.Video {
	var video entity.Video
	err := context.BindJSON(&video)
	if err != nil {
		return entity.Video{}
	}
	c.service.Save(video)
	return video
}

func New(videoService service.VideoService) VideoController {
	return &controller{
		service: videoService,
	}
}
