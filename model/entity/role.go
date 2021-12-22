package entity

import (
	"go-roles-management/app/database"

	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name            string
	Label           string
	Description     string
	Users           []User
	RolesPermission []RolePermission
}

func RoleMigrate() {
	db := database.NewDB()
	db.AutoMigrate(&Role{})
}
