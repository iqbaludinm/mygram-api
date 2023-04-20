package models

type Comment struct {
	GormModel
	UserID  uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PhotoID uint   `json:"photo_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;index"`
	User    *User  `json:"user_data"`
	Photo   *Photo `json:"photo_data"`
	Message string `json:"message" gorm:"not null"`
}

type InsertComment struct {
	UserID  uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" example:"1"`
	PhotoID uint   `json:"photo_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;index" example:"1"`
	Message string `json:"message" gorm:"not null" example:"Nice Info!"`
}

type UpdateComment struct {
	UserID  uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" example:"1"`
	PhotoID uint   `json:"photo_id" binding:"omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;index" example:"1"`
	Message string `json:"message" binding:"omitempty" example:"I Wanna Hire You. Please, check your DM!"`
}