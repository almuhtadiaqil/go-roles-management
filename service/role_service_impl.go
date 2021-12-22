package service

import (
	"go-roles-management/exception"
	"go-roles-management/helper"
	"go-roles-management/model/entity"
	"go-roles-management/repository"
	"strconv"

	"go-roles-management/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewRoleService(rr repository.RoleRepository, db *gorm.DB, v *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: rr,
		DB:             db,
		Validate:       v,
	}
}

func (service *RoleServiceImpl) Save(ctx gin.Context, request web.RoleCreateRequest) web.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	role := entity.Role{
		Name:        request.Name,
		Label:       request.Label,
		Description: request.Description,
	}

	role = service.RoleRepository.Save(ctx, tx, role)

	return helper.ToRoleResponse(role)
}
func (service *RoleServiceImpl) Update(ctx gin.Context, request web.RoleUpdateRequest) web.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	roleId := ctx.Param("roleId")
	id, _ := strconv.Atoi(roleId)

	role := entity.Role{
		Name:        request.Name,
		Label:       request.Label,
		Description: request.Description,
	}

	role = service.RoleRepository.Update(ctx, tx, role, uint(id))

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Delete(ctx gin.Context, roleId uint) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.RoleRepository.Delete(ctx, tx, role)
}
func (service *RoleServiceImpl) FindById(ctx gin.Context, roleId uint) web.RoleResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindAll(ctx gin.Context) []web.RoleResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	roles := service.RoleRepository.FindAll(ctx, tx)

	return helper.ToRoleResponses(roles)
}
