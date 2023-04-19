package models

type Photo struct {
	GormModel
	Title    string    `json:"title" gorm:"not null"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photo_url" binding:"required" gorm:"not null"`
	UserID   uint      `json:"user_id" gorm:"constraint"`
	User     *User     `json:"user_data"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PhotoID;constraint"`
}

type InsertPhoto struct {
	Title    string `json:"title" binding:"required,min=3"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint   `json:"user_id"`
}

type UpdatePhoto struct {
	Title    string `json:"title" binding:"omitempty,min=3"`
	Caption  string `json:"caption" binding:"omitempty"`
	PhotoUrl string `json:"photo_url" binding:"omitempty"`
	UserID   uint   `json:"user_id"`
}
