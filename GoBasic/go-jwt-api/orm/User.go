package orm

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avater   string
}
