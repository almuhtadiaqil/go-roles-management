package main

import (
	"go-roles-management/app/database"
	"go-roles-management/controller"
	"go-roles-management/helper"
	"go-roles-management/model/entity"
	"go-roles-management/repository"
	"go-roles-management/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	entity.RoleMigrate()
	entity.PermissionMigrate()
	entity.UserMigrate()
	entity.RolePermissionMigrate()
}

func main() {
	db := database.NewDB()
	// Migrate(db)
	validate := validator.New()
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	roles := v1.Group("/roles")

	roles.GET("", roleController.FindAll)
	roles.GET("/:roleId", roleController.FindById)
	roles.POST("", roleController.Save)
	roles.PUT("/:roleId", roleController.Update)
	roles.DELETE("/:roleId", roleController.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
