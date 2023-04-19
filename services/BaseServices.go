package services

import "github.com/iqbaludinm/mygram-api/repositories"

type BaseService struct {
	repo repositories.RepoInterface
}

type ServiceInterface interface {
	UserService
	PhotoService
	SocialMediaService
	CommentService
}

func NewService(repo repositories.RepoInterface) ServiceInterface {
	return &BaseService{repo: repo}
}