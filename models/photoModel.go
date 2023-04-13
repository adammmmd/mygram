package models

import (
	"github.com/asaskevich/govalidator"
	
	"time"
	"gorm.io/gorm"
)

type Photo struct {
	ID 			uint		`json:"id" gorm:"primaryKey"`
	Title 		string		`json:"title" gorm:"not null" form:"title" valid:"required"`
	Caption 	string		`json:"caption" form:"caption"`
	PhotoURL 	string		`json:"photo_url" gorm:"not null" form:"photo_url" valid:"required"`
	UserID		uint		`json:"user_id" `
	CreatedAt 	*time.Time
	UpdatedAt 	*time.Time 
	Comments	[]Comment	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}