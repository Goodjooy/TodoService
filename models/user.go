package models

import (
	"todo-web/server/server"

	"github.com/jinzhu/gorm")


type UserModel struct {
	gorm.Model

	Name         string `gorm:"type:varchar(32);not null" admin:"name:用户名;type:text"`
	EmailAddress string `gorm:"type:varchar(128)" admin:"name:邮箱;type:email"`

	PassWord string `gorm:"type:varchar(32);not null" admin:"name:用户密码哈希;256G;type:password"`

	Todos []Todo `gorm:"one-to-many"`
}

func FromUserClaims(user server.UserClaims)UserModel{
	return UserModel{
		Model: gorm.Model{ID: user.ID},
		EmailAddress: user.Name,
		PassWord: user.Password,
	}
}
func FromUser(user UserModel)*server.UserClaims{
	return &server.UserClaims{
		ID: user.ID,
		Name: user.EmailAddress,
		Password: user.PassWord,
	}
}