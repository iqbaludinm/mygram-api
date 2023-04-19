package repositories

import (
	"errors"

	"github.com/iqbaludinm/mygram-api/models"
	"gorm.io/gorm"
)

// interface
type PhotoRepository interface {
	CreatePhoto(photo models.Photo) (res models.Photo, err error)
	GetPhotos() (photos []models.Photo, err error)
	GetPhotoById(id int64) (photo models.Photo, err error)
	UpdatePhoto(id int64, photo models.Photo) (res models.Photo, err error)
	DeletePhoto(id int64, userId uint) (res models.Photo, err error)
}

func (r BaseRepository) CreatePhoto(photo models.Photo) (res models.Photo, err error) {
	err = r.gorm.Create(&photo).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r BaseRepository) GetPhotos() (photos []models.Photo, err error) {
	err = r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Select("id, username, email, age")
    }).Preload("Comments").Find(&photos).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return photos, nil
		}
		return nil, err
	}

	return photos, err
}

func (r BaseRepository) GetPhotoById(id int64) (photo models.Photo, err error) {
	err = r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Select("id, username, email, age")
    }).Preload("Comments").First(&photo, id).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r BaseRepository) UpdatePhoto(id int64, photo models.Photo) (res models.Photo, err error) {
	err = r.gorm.Model(&photo).Where("id = ?", id).First(&res).Error
	if err != nil {
		return res, err
	}
	
	// validate id user logged in with database
	if res.UserID != photo.UserID {
		return res, errors.New("unauthorized")
	}

	err = r.gorm.Model(&res).Updates(photo).Error
	
	if err != nil {
		return res, err
	}
	
	return res, nil
}

func (r BaseRepository) DeletePhoto(id int64, userId uint) (res models.Photo, err error) {
	photo := models.Photo{}
	err = r.gorm.First(&photo, id).Scan(&res).Error

	if err != nil {
		return res, err
	}
	
	if res.UserID != userId {
		return res, errors.New("unauthorized")
	}

	err = r.gorm.Delete(&photo, id, userId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}
