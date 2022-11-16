package auth

import "gorm.io/gorm"

type Account struct {
	gorm.Model `json:"-"`
	UserName   string `gorm:"username" json:"username"`
	Password   string `gorm:"password" json:"password"`
}
