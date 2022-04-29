package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaUrl string `gorm:"not null" form:"social_media_url" json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	User           User   `json:"user"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(s.Name)) == 0 {
		err = errors.New("name is required")
		return
	}
	if len(strings.TrimSpace(s.SocialMediaUrl)) == 0 {
		err = errors.New("url is required")
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(s.Name)) == 0 {
		err = errors.New("name is required")
		return
	}
	if len(strings.TrimSpace(s.SocialMediaUrl)) == 0 {
		err = errors.New("url is required")
		return
	}

	err = nil
	return
}
