package services

import (
	"github.com/iqbaludinm/mygram-api/models"
)

// interface
type PhotoService interface {
	CreatePhoto(photo models.InsertPhoto) (models.Photo, error)
	GetPhotos() ([]models.Photo, error)
	GetPhotoById(id int64) (models.Photo, error)
	UpdatePhoto(id int64, photo models.UpdatePhoto) (models.Photo, error)
	DeletePhoto(id int64, userId uint) (models.Photo, error)
}

func (s *BaseService) CreatePhoto(photo models.InsertPhoto) (models.Photo, error) {
	photos := models.Photo{
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID: photo.UserID,
	}
	return s.repo.CreatePhoto(photos)
}

func (s *BaseService) GetPhotos() ([]models.Photo, error) {
	return s.repo.GetPhotos()
}

func (s *BaseService) GetPhotoById(id int64) (models.Photo, error) {
	return s.repo.GetPhotoById(id)
}

func (s *BaseService) UpdatePhoto(id int64, photo models.UpdatePhoto) (models.Photo, error) {
	photos := models.Photo{
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID: photo.UserID,
	}
	return s.repo.UpdatePhoto(id, photos)
}

func (s *BaseService) DeletePhoto(id int64, userId uint) (models.Photo, error) {
	return s.repo.DeletePhoto(id, userId)
}