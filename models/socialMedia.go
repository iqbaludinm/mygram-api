package models

type SocialMedia struct {
	GormModel
	Name           string `json:"name" gorm:"not null"`
	SocialMediaUrl string `json:"social_media_url" binding:"required" gorm:"not null"`
	UserID         uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User *User `json:"user_data"`
}

type InsertSocialMedia struct {
	Name           string `json:"name" gorm:"not null"`
	SocialMediaUrl string `json:"social_media_url" binding:"required" gorm:"not null"`
	UserID         uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UpdateSocialMedia struct {
	Name           string `json:"name" binding:"omitempty"`
	SocialMediaUrl string `json:"social_media_url" binding:"omitempty"`
	UserID         uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
