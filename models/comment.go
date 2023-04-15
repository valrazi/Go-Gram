package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint   `json:"photo_id"`
	Message string `json:"msg" form:"msg" valid:"required~Comment message is required"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	err = nil
	return
}
