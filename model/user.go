// File:    user
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/30 22:16
// DESC:    blog用户操作

package model

import "time"

type User struct {
	Id        int       `gorm:"id"`
	UserName  string    `gorm:"username"`
	Password  string    `gorm:"password"`
	RoleId    int       `gorm:"role_id"`
	Avatar    string    `gorm:"avatar"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (User)TableName() string {
	return `user`
}

// GetUserInfoRow get the one user information by username or other condition.
func GetUserInfoRow(filter interface{}) (row *User, err error) {
	err = db.Where(filter).Find(&row).Error
	return
}

//
func CreateNewUser(newUser *User) (userId int){
	err := db.Create(&newUser).Error
	if err != nil {
	}
	return newUser.Id
}
