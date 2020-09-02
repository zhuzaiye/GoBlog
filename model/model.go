// File:    model
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/30 22:17
// DESC:

package model

import (
	"GoBlog/utils/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db *gorm.DB
)

func InitMysqlDB() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.UserName, setting.Password, setting.Host, setting.Port, setting.DataBase,
	))
	if err != nil {
		fmt.Println("link mysql database failed")
	}
	//设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(30)
	//设置打开数据库连接的最大数量
	db.DB().SetMaxOpenConns(100)
	//设置了连接可复用的最大时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//_ = db.Close()
}
