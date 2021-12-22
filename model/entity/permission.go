package entity

import (
	"go-roles-management/app/database"

	"github.com/jinzhu/gorm"
)

type Permission struct {
	gorm.Model
	Title           string
	Description     string
	RolesPermission []RolePermission
}

func PermissionMigrate() {
	db := database.NewDB()
	db.AutoMigrate(&Permission{})
}
