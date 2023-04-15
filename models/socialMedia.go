package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required~Social Media Name is required"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         uint
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(sm)

	if errValidate != nil {
		err = errValidate
		return
	}

	err = nil
	return
}

func (sm *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValidate := govalidator.ValidateStruct(sm)

	if errValidate != nil {
		err = errValidate
		return
	}

	err = nil
	return
}
