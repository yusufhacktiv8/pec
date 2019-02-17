package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"pecrsuh.id/pec/models"
)

func TestFindRoles(t *testing.T) {
	a := GetAppTest()
	gf := a.GoFight

	a.DB.Unscoped().Delete(&models.Role{})

	gf.GET("/api/roles/").
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"message\":\"No count parameter\",\"status\":400}", r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	gf.GET("/api/roles/?count=10").
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"message\":\"No start parameter\",\"status\":400}", r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	newRole := models.Role{Code: "ADMIN", Name: "Admin"}
	a.DB.Create(&newRole)

	gf.GET("/api/roles/?count=10&start=0").
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			var objMap map[string]*json.RawMessage
			json.Unmarshal([]byte(r.Body.String()), &objMap)

			var roles []models.Role
			json.Unmarshal(*objMap["data"], &roles)

			if len(roles) > 0 {
				role := roles[0]

				assert.Equal(t, "ADMIN", role.Code)
				assert.Equal(t, http.StatusOK, r.Code)
			}

		})
}

func TestCreateRole(t *testing.T) {
	a := GetAppTest()
	gf := a.GoFight

	a.DB.Unscoped().Delete(&models.Role{})

	gf.POST("/api/roles/").
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)

			var objMap map[string]*json.RawMessage
			json.Unmarshal([]byte(r.Body.String()), &objMap)

			var message string
			json.Unmarshal(*objMap["message"], &message)

			assert.Equal(t, "Code or Name is empty", message)

		})

	gf.POST("/api/roles/").
		SetJSON(gofight.D{
			"code": "ADMIN",
			"name": "Admin",
		}).
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)

			var objMap map[string]*json.RawMessage
			json.Unmarshal([]byte(r.Body.String()), &objMap)

			if http.StatusCreated == r.Code {
				var roleID uint
				json.Unmarshal(*objMap["resourceId"], &roleID)
				assert.NotEqual(t, 0, roleID)
			}
		})

	gf.POST("/api/roles/").
		SetJSON(gofight.D{
			"code": "ADMIN",
			"name": "Admin",
		}).
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)

			var objMap map[string]*json.RawMessage
			json.Unmarshal([]byte(r.Body.String()), &objMap)

			var message string
			json.Unmarshal(*objMap["message"], &message)

			assert.Equal(t, "Code is not unique", message)
		})
}

func TestUpdateRole(t *testing.T) {
	a := GetAppTest()
	gf := a.GoFight

	a.DB.Unscoped().Delete(&models.Role{})

	newRole := models.Role{Code: "ADMIN", Name: "Admin"}
	a.DB.Create(&newRole)

	gf.PUT("/api/roles/0").
		SetJSON(gofight.D{}).
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)

			if http.StatusNotFound == r.Code {
				var objMap map[string]*json.RawMessage
				json.Unmarshal([]byte(r.Body.String()), &objMap)

				var message string
				json.Unmarshal(*objMap["message"], &message)

				assert.Equal(t, "Role not found", message)
			}
		})

	gf.PUT("/api/roles/"+strconv.Itoa(int(newRole.ID))).
		SetJSON(gofight.D{
			"code": "",
			"name": "",
		}).
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)

			if http.StatusBadRequest == r.Code {
				var objMap map[string]*json.RawMessage
				json.Unmarshal([]byte(r.Body.String()), &objMap)

				var message string
				json.Unmarshal(*objMap["message"], &message)

				assert.Equal(t, "Code or Name is empty", message)
			}
		})

	gf.PUT("/api/roles/"+strconv.Itoa(int(newRole.ID))).
		SetJSON(gofight.D{
			"code": "USER",
			"name": "User",
		}).
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			if http.StatusOK == r.Code {
				var role models.Role
				a.DB.First(&role)

				assert.Equal(t, "USER", role.Code)
				assert.Equal(t, "User", role.Name)

			}
		})
}

func TestDeleteRole(t *testing.T) {
	a := GetAppTest()
	gf := a.GoFight

	a.DB.Unscoped().Delete(&models.Role{})

	newRole := models.Role{Code: "ADMIN", Name: "Admin"}
	a.DB.Create(&newRole)

	gf.DELETE("/api/roles/0").
		SetJSON(gofight.D{}).
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)

			if http.StatusNotFound == r.Code {
				var objMap map[string]*json.RawMessage
				json.Unmarshal([]byte(r.Body.String()), &objMap)

				var message string
				json.Unmarshal(*objMap["message"], &message)

				assert.Equal(t, "Role not found", message)
			}
		})

	gf.DELETE(fmt.Sprintf("/api/roles/%d", newRole.ID)).
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
