package models

import (
	"go-final-task/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex"json:"username" valid:"required""`
	Email    string `gorm:"not null;uniqueIndex"json:"email" valid:"required,email""`
	Password string `gorm:"not null"json:"password" valid:"required,minstringlength(6)"`
	Age      uint   `gorm:"not null"json:"age" valid:"required,numeric"`
	// Photo []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Photo"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(u)

	if errValidate != nil {
		err = errValidate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
