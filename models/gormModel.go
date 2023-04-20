package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint      `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
}

func (model *GormModel) BeforeCreate(tx *gorm.DB) (err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("Error LoadLocation: ", err)
		return err
	}
	model.CreatedAt = time.Now().In(location)
	model.UpdatedAt = time.Now().In(location)
	return nil
}

func (model *GormModel) BeforeUpdate(tx *gorm.DB) (err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("Error LoadLocation: ", err)
		return err
	}
	model.UpdatedAt = time.Now().In(location)
	return nil
}
