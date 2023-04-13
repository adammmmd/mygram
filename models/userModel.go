package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"Project/helpers"
	"time"
)

type User struct {
	ID 			uint			`json:"id" gorm:"primaryKey"`
	Username 	string			`json:"username" gorm:"not null;uniqueIndex" form:"username" valid:"required"`
	Email 		string			`json:"email" gorm:"not null;uniqueIndex" form:"email" valid:"required,email"`
	Password 	string			`json:"password" gorm:"not null" form:"password" valid:"required,minstringlength(6)"`
	Age 		uint			`json:"age" gorm:"not null" form:"age" valid:"required,range(8|100)"`
	CreatedAt	*time.Time 	
	UpdatedAt	*time.Time
	Photos		[]Photo			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Socmed		[]SocialMedia	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socmed"`
	Comments	[]Comment		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
