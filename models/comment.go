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
	UserID  uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PhotoID uint   `json:"photo_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;index"`
	Message string `json:"message" gorm:"not null"`
}

type UpdateComment struct {
	UserID  uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PhotoID uint   `json:"photo_id" binding:"omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;index"`
	Message string `json:"message" binding:"omitempty"`
}