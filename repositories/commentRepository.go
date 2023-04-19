package repositories

import (
	"errors"

	"github.com/iqbaludinm/mygram-api/models"
	"gorm.io/gorm"
)

// interface
type CommentRepository interface {
	CreateComment(comment models.Comment) (res models.Comment, err error)
	GetComments(photoId uint) (comments []models.Comment, err error)
	GetCommentById(id int64) (comment models.Comment, err error)
	UpdateComment(id int64, comment models.Comment) (res models.Comment, err error)
	DeleteComment(id int64, userId uint) (res models.Comment, err error)
}

func (r BaseRepository) CreateComment(comment models.Comment) (res models.Comment, err error) {
	photo := models.Photo{}
	if err := r.gorm.First(&photo, comment.PhotoID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.New("photo not found")
		}
		return res, err
	}

	user := models.User{}
	if err := r.gorm.First(&user, comment.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.New("user not found")
		}
		return res, err
	}

	err = r.gorm.Create(&comment).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r BaseRepository) GetComments(photoId uint) (comments []models.Comment, err error) {
	err = r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age")
	}).Preload("Photo").Where("photo_id = ?", photoId).Find(&comments).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return comments, nil
		}
		return nil, err
	}

	return comments, err
}

func (r BaseRepository) GetCommentById(id int64) (comment models.Comment, err error) {
	err = r.gorm.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age")
	}).Preload("Photo").First(&comment, id).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r BaseRepository) UpdateComment(id int64, comment models.Comment) (res models.Comment, err error) {
	err = r.gorm.Model(&comment).Where("id = ?", id).First(&res).Error
	if err != nil {
		return res, err
	}

	if res.UserID != comment.UserID {
		return res, errors.New("unauthorized")
	}

	err = r.gorm.Model(&res).Updates(comment).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r BaseRepository) DeleteComment(id int64, userId uint) (res models.Comment, err error) {
	comment := models.Comment{}
	err = r.gorm.First(&comment, id).Scan(&res).Error

	if err != nil {
		return res, err
	}

	if res.UserID != userId {
		return res, errors.New("unauthorized")
	}

	err = r.gorm.Delete(&comment, id, userId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}
