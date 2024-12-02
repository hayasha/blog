package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID       uint
	Content  string
	PostID   uint
	Nickname string
}
