package services

import "github.com/iqbaludinm/mygram-api/models"

// interface
type SocialMediaService interface {
	CreateSocmed(socmed models.InsertSocialMedia) (models.SocialMedia,error)
	GetSocmeds() ([]models.SocialMedia, error)
	GetSocmedById(id int64) (models.SocialMedia, error)
	UpdateSocmed(id int64, socmed models.UpdateSocialMedia) (models.SocialMedia, error)
	DeleteSocmed(id int64, userId uint) (models.SocialMedia, error)
}

func (s *BaseService) CreateSocmed(socmed models.InsertSocialMedia) (models.SocialMedia, error) {
	socmeds := models.SocialMedia{
		Name: socmed.Name,
		SocialMediaUrl: socmed.SocialMediaUrl,
		UserID: socmed.UserID,
	}
	return s.repo.CreateSocmed(socmeds)
}

func (s *BaseService) GetSocmeds() ([]models.SocialMedia, error) {
	return s.repo.GetSocmeds()
}

func (s *BaseService) GetSocmedById(id int64) (models.SocialMedia, error) {
	return s.repo.GetSocmedById(id)
}

func (s *BaseService) UpdateSocmed(id int64,socmed models.UpdateSocialMedia) (models.SocialMedia, error) {
	socmeds := models.SocialMedia{
		Name: socmed.Name,
		SocialMediaUrl: socmed.SocialMediaUrl,
		UserID: socmed.UserID,
	}
	return s.repo.UpdateSocmed(id, socmeds)
}

func (s *BaseService) DeleteSocmed(id int64, userId uint) (models.SocialMedia, error) {
	return s.repo.DeleteSocmed(id, userId)
}