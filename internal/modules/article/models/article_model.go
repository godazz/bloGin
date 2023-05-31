package models

import (
	"github.com/godazz/bloGin/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:191"`
	Content string `gorm:"varchar:191"`
	UserID  uint
	User    models.User
}
