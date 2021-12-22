package entity

import (
	"go-roles-management/app/database"

	"github.com/jinzhu/gorm"
)

type RolePermission struct {
	gorm.Model
	RoleId       uint
	PermissionId uint
}

func RolePermissionMigrate() {
	db := database.NewDB()
	db.AutoMigrate(&RolePermission{})
}
