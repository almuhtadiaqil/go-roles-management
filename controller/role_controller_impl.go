package controller

import (
	"go-roles-management/helper"
	"go-roles-management/model/web"
	"go-roles-management/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) Save(ctx *gin.Context) {
	roleCreateRequest := web.RoleCreateRequest{}
	helper.ReadFromRequestBody(ctx, &roleCreateRequest)

	roleResponse := controller.RoleService.Save(*ctx, roleCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    roleResponse,
	}
	helper.WriteToResponseBody(ctx, webResponse)
}
func (controller *RoleControllerImpl) Update(ctx *gin.Context) {
	roleUpdateRequest := web.RoleUpdateRequest{}
	helper.ReadFromRequestBody(ctx, &roleUpdateRequest)

	roleId := ctx.Param("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)
	roleUpdateRequest.Id = uint(id)

	roleResponse := controller.RoleService.Update(*ctx, roleUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    roleResponse,
	}
	helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *RoleControllerImpl) Delete(ctx *gin.Context) {
	roleId := ctx.Param("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	controller.RoleService.Delete(*ctx, uint(id))
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
	}
	helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *RoleControllerImpl) FindById(ctx *gin.Context) {
	roleId := ctx.Param("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	roleResponse := controller.RoleService.FindById(*ctx, uint(id))
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    roleResponse,
	}
	helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *RoleControllerImpl) FindAll(ctx *gin.Context) {
	roleResponses := controller.RoleService.FindAll(*ctx)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    roleResponses,
	}
	helper.WriteToResponseBody(ctx, webResponse)
}
