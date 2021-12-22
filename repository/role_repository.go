package repository

import (
	"go-roles-management/model/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(ctx gin.Context, tx *gorm.DB, role entity.Role) entity.Role
	Update(ctx gin.Context, tx *gorm.DB, role entity.Role, roleId uint) entity.Role
	Delete(ctx gin.Context, tx *gorm.DB, role entity.Role)
	FindById(ctx gin.Context, tx *gorm.DB, roleId uint) (entity.Role, error)
	FindAll(ctx gin.Context, tx *gorm.DB) []entity.Role
}
