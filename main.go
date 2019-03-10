package main

import (
	"fmt"
	"os"

	"pecrsuh.id/pec/application"
	"pecrsuh.id/pec/security"
)

func main() {
	a := application.App{}
	fmt.Printf("Running PEC Server...\n")
	a.InitDatabase(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	a.InitRoutes()
	setControllers(&a)
	a.Router.Run(":8080")
}

func setControllers(a *application.App) {
	roleController := security.RoleController{DB: a.DB}

	v1 := a.Router.Group("/api/roles")
	{
		v1.POST("/", roleController.CreateRole)
		v1.GET("/", roleController.FindRoles)
		v1.PUT("/:id", roleController.UpdateRole)
		v1.DELETE("/:id", roleController.DeleteRole)
	}
}
