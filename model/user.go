package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}
