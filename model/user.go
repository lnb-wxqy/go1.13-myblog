package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;unique"`
	Telephone string `gorm:"type:varchar(11)"`
	Password  string `gorm:"type:varchar(110);not null"`
}
