package service

import "gin_server/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct { // Will implement VideoService
	videos []entity.Video // Have a variable is videos
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

func New() VideoService {
	return &videoService{}
}
