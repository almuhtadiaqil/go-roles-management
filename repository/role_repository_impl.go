package repository

import (
	"go-roles-management/model/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Save(ctx gin.Context, tx *gorm.DB, role entity.Role) entity.Role {
	tx.Create(&role)
	return role
}
func (repository *RoleRepositoryImpl) Update(ctx gin.Context, tx *gorm.DB, role entity.Role, roleId uint) entity.Role {
	tx.Where("id", roleId)
	tx.Save(&role)
	return role
}
func (repository *RoleRepositoryImpl) Delete(ctx gin.Context, tx *gorm.DB, role entity.Role) {
	tx.Delete(&role)
}

func (repository *RoleRepositoryImpl) FindById(ctx gin.Context, tx *gorm.DB, roleId uint) (entity.Role, error) {
	role := entity.Role{}
	result := tx.Find(&role, roleId).Where("deleted_at = ?", nil)
	return role, result.Error
}
func (repository *RoleRepositoryImpl) FindAll(ctx gin.Context, tx *gorm.DB) []entity.Role {
	role := []entity.Role{}
	tx.Find(&role)
	return role
}
