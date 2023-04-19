package repositories

import "gorm.io/gorm"

type BaseRepository struct {
	gorm *gorm.DB
}

type RepoInterface interface {
	
	UserRepository
	PhotoRepository
	SocialMediaRepository
	CommentRepository
}

// constructor
func NewRepo(gorm *gorm.DB) *BaseRepository {
	return &BaseRepository{gorm: gorm}
}