package entity

import (
	"go-roles-management/app/database"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	RoleId   uint
	Name     string
	Email    string
	Username string
	Password string
	Mobile   string
}

func UserMigrate() {
	db := database.NewDB()
	db.AutoMigrate(&User{})
}
