// File:    article_category
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/30 22:17
// DESC:

package model

import "github.com/jinzhu/gorm"

type ArticleCategory struct {
	gorm.Model
	Name string
}