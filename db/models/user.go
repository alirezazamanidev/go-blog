package models

import "gorm.io/gorm"


type User struct {
	gorm.Model

	Phone        string  `gorm:"not null" json:"phone"`
	Username     *string `gorm:"unique" json:"username"`
	Password     string  `gorm:"not null" json:"-"`
	PhoneVerify  bool    `gorm:"default:false" json:"phone_verify"`
}
