package services

import "github.com/iqbaludinm/mygram-api/models"

// interface
type CommentService interface {
	CreateComment(comment models.InsertComment) (models.Comment, error)
	GetComments(photoId uint) ([]models.Comment, error)
	GetCommentById(id int64) (models.Comment, error)
	UpdateComment(id int64, comment models.UpdateComment) (models.Comment, error)
	DeleteComment(id int64, userId uint) (models.Comment, error)
}

func (s *BaseService) CreateComment(comment models.InsertComment) (models.Comment, error) {
	comments := models.Comment{
		UserID: comment.UserID,
		PhotoID: comment.PhotoID,
		Message: comment.Message,
	}
	return s.repo.CreateComment(comments)
}

func (s *BaseService) GetComments(photoId uint) ([]models.Comment, error) {
	return s.repo.GetComments(photoId)
}

func (s *BaseService) GetCommentById(id int64) (models.Comment, error) {
	return s.repo.GetCommentById(id)
}

func (s *BaseService) UpdateComment(id int64, comment models.UpdateComment) (models.Comment, error) {
	comments := models.Comment{
		PhotoID: comment.PhotoID,
		UserID: comment.UserID,
		Message: comment.Message,
	}
	return s.repo.UpdateComment(id, comments)
}

func (s *BaseService) DeleteComment(id int64, userId uint) (models.Comment, error) {
	return s.repo.DeleteComment(id, userId)
}