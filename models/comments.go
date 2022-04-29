package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Comments struct {
	GormModel
	UserID  uint   `json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message"`
	User    User   `json:"user"`
	Photo   Photo  `json:"photo"`
}

func (c *Comments) BeforeCreate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(c.Message)) == 0 {
		err = errors.New("message is required")
		return
	}
	err = nil
	return
}
