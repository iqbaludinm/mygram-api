package repositories

import (
	"errors"

	"github.com/iqbaludinm/mygram-api/models"
	"gorm.io/gorm"
)

// interface
type SocialMediaRepository interface {
	CreateSocmed(socmed models.SocialMedia) (res models.SocialMedia, err error)
	GetSocmeds() (socmeds []models.SocialMedia, err error)
	GetSocmedById(id int64) (socmed models.SocialMedia, err error)
	UpdateSocmed(id int64, socmed models.SocialMedia) (res models.SocialMedia, err error)
	DeleteSocmed(id int64, userId uint) (res models.SocialMedia, err error)
}

func (r BaseRepository) CreateSocmed(socmed models.SocialMedia) (res models.SocialMedia, err error) {
	err = r.gorm.Create(&socmed).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r BaseRepository) GetSocmeds() (socmeds []models.SocialMedia, err error) {
	err = r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age")
	}).Find(&socmeds).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return socmeds, nil
		}
		return nil, err
	}

	return socmeds, err
}

func (r BaseRepository) GetSocmedById(id int64) (socmed models.SocialMedia, err error) {
	err = r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age")
	}).First(&socmed, id).Error
	if err != nil {
		return socmed, err
	}

	return socmed, nil
}

func (r BaseRepository) UpdateSocmed(id int64,socmed models.SocialMedia) (res models.SocialMedia, err error) {
	err = r.gorm.Model(&socmed).Where("id = ?", id).First(&res).Error

	if err != nil {
		return res, err
	}

	// validate id user logged in with database
	if res.UserID != socmed.UserID {
		return res, errors.New("unauthorized")
	}

	err = r.gorm.Model(&res).Updates(socmed).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r BaseRepository) DeleteSocmed(id int64, userId uint) (res models.SocialMedia, err error) {
	socmed := models.SocialMedia{}
	err = r.gorm.First(&socmed, id).Scan(&res).Error

	if err != nil {
		return res, err
	}
	
	if res.UserID != userId {
		return res, errors.New("unauthorized")
	}

	err = r.gorm.Delete(&socmed, id, userId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}
