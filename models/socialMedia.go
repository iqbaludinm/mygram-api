package models

type SocialMedia struct {
	GormModel
	Name           string `json:"name" gorm:"not null"`
	SocialMediaUrl string `json:"social_media_url" binding:"required" gorm:"not null"`
	UserID         uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User           *User  `json:"user_data"`
}

type InsertSocialMedia struct {
	Name           string `json:"name" gorm:"not null" example:"Linkedin Muhammad Iqbaludin Zaky"`
	SocialMediaUrl string `json:"social_media_url" binding:"required" gorm:"not null" example:"https://www.linkedin.com/in/iqbaludinm"`
	UserID         uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" example:"1"`
}

type UpdateSocialMedia struct {
	Name           string `json:"name" binding:"omitempty"`
	SocialMediaUrl string `json:"social_media_url" binding:"omitempty" example:"https://tinyurl.com/linkedin-iqbal"`
	UserID         uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" example:"1"`
}
