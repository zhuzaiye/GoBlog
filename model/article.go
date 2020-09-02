// File:    article
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/30 22:16
// DESC:    article table

package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title       string `gorm:"title"`
	CategoryId  int    `gorm:"category_id"`
	Description string `gorm:"description"`
	Content     string `gorm:"content"`
	Image       string `gorm:"image"`
}
