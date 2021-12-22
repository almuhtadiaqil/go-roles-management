package helper

import (
	"go-roles-management/model/entity"
	"go-roles-management/model/web"
)

func ToRoleResponse(role entity.Role) web.RoleResponse {
	return web.RoleResponse{
		Id:          int(role.ID),
		Name:        role.Name,
		Label:       role.Label,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

func ToRoleResponses(roles []entity.Role) []web.RoleResponse {
	var roleResponses []web.RoleResponse

	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}
