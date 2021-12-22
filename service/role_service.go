package service

import (
	"go-roles-management/model/web"

	"github.com/gin-gonic/gin"
)

type RoleService interface {
	Save(ctx gin.Context, request web.RoleCreateRequest) web.RoleResponse
	Update(ctx gin.Context, request web.RoleUpdateRequest) web.RoleResponse
	Delete(ctx gin.Context, roleId uint)
	FindById(ctx gin.Context, roleId uint) web.RoleResponse
	FindAll(ctx gin.Context) []web.RoleResponse
}
