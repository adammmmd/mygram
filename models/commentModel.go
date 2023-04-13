package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Comment struct {
	ID 			uint		`json:"id" gorm:"not null"`
	UserID		uint		`json:"user_id"`
	PhotoID		uint		`json:"photo_id" gorm:"not null" form:"photo_id" valid:"required"`
	Message 	string		`json:"message" gorm:"not null" form:"message" valid:"required"`
	CreatedAt 	*time.Time
	UpdateAt 	*time.Time
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}