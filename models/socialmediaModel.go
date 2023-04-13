package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/asaskevich/govalidator"
)

type SocialMedia struct {
	ID 				uint		`json:"id" gorm:"not null"`
	Name 			string		`json:"name" gorm:"not null" form:"name" valid:"required"`		
	SocialMediaURL 	string		`json:"socmed_url" gorm:"not null" form:"socmed_url" valid:"required"`
	UserID 			uint		`json:"user_id"`
	CreatedAt 		*time.Time
	UpdatedAt 		*time.Time
} 

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}