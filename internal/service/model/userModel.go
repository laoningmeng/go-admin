package model

import (
	"github.com/laoningmeng/go-admin/internal/service/global"
)

type UserStatus int

const (
	PreEmployment UserStatus = 0 // 预入职
	Employed      UserStatus = 1 // 在职
	Left          UserStatus = 2 // 离职
)

type User struct {
	BaseItem
	Name         string     `json:"name"`
	Account      string     `json:"account"`
	Email        string     `json:"email"`
	Password     string     `json:"password" gorm:"not null;"`
	Avatar       string     `json:"avatar"`
	Introduction string     `json:"introduction"`
	Status       UserStatus `json:"status" gorm:"default:0"`
	Role         Role       `json:"role" gorm:"foreignKey:ID"`
}

func (u User) FindOne() (User, error) {
	res := User{}
	result := global.DB.Where(u).Find(&res)
	if result.Error != nil {
		return User{}, result.Error
	}
	return res, nil
}
