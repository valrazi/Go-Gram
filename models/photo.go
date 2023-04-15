package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" valid:"required~Photo title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	err = nil
	return
}
