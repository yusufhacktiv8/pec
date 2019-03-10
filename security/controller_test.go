package security

import (
	"os"
	"testing"

	"pecrsuh.id/pec/application"

	"github.com/appleboy/gofight"
)

var a application.App

func TestMain(m *testing.M) {
	a = application.App{}
	a.InitDatabase("myyusuf", "", "pec")
	a.InitRoutes()
	setControllers(&a)
	defer a.DB.Close()

	a.GoFight = gofight.New()

	theRun := m.Run()

	os.Exit(theRun)
}

func GetAppTest() *application.App {
	return &a
}

func setControllers(a *application.App) {
	roleController := RoleController{DB: a.DB}

	v1 := a.Router.Group("/api/roles")
	{
		v1.POST("/", roleController.CreateRole)
		v1.GET("/", roleController.FindRoles)
		v1.PUT("/:id", roleController.UpdateRole)
		v1.DELETE("/:id", roleController.DeleteRole)
	}
}
