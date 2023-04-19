package models

import (
	"github.com/iqbaludinm/mygram-api/helpers"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string       `json:"username" form:"username" binding:"required" gorm:"not null;uniqueIndex"`
	Email       string       `json:"email" form:"email" binding:"required,email" gorm:"not null;uniqueIndex"`
	Password    string       `json:"password" form:"password" binding:"min=6" gorm:"not null"`
	Age         int          `json:"age" form:"age" binding:"required,gt=8" gorm:"not null"`
	Photos      []Photo      `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;reference:ID"`
	Comments    []Comment    `json:"comments" gorm:"foreignKey:UserID"`
	SocialMedia *SocialMedia `json:"social_media,omitempty" gorm:"foreignKey:UserID"`
}

type LoginUser struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return err
}
