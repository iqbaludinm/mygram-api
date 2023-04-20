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
	Title    string `json:"title" binding:"required,min=3" example:"My Bootcamp Journey in Hacktiv8"`
	Caption  string `json:"caption" example:"This is my second bootcamp experience"`
	PhotoUrl string `json:"photo_url" binding:"required" example:"https://images/image-1.jpg"`
	UserID   uint   `json:"user_id" example:"1"`
}

type UpdatePhoto struct {
	Title    string `json:"title" binding:"omitempty,min=3" example:"Hacktiv8 Golang Class"`
	Caption  string `json:"caption" binding:"omitempty" example:"Hoho, now i am a Golang Developer"`
	PhotoUrl string `json:"photo_url" binding:"omitempty"`
	UserID   uint   `json:"user_id" example:"1"`
}
