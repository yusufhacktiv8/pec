package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"pecrsuh.id/pec/models"
)

// RoleController for role feature
type RoleController struct {
	DB *gorm.DB
}

// FindRoles find roles within range defined by start and count query parameters
// Use the searchText query parameter to filter code or name (incasesensitive)
func (a *RoleController) FindRoles(c *gin.Context) {
	countStr, ok := c.GetQuery("count")
	if !ok {
		SendBadRequest(c, "No count parameter")
		return
	}

	startStr, ok := c.GetQuery("start")
	if !ok {
		SendBadRequest(c, "No start parameter")
		return
	}

	count, _ := strconv.Atoi(countStr)
	start, _ := strconv.Atoi(startStr)
	searchText, _ := c.GetQuery("searchText")
	searchText = strings.ToLower(searchText)

	var roles []models.Role
	a.DB.Where("lower(code) LIKE ?", "%"+searchText+"%").Or("lower(name) LIKE ?", "%"+searchText+"%").Offset(start).Limit(count).Find(&roles)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": roles})

}

// CreateRole create Role with code and name as parameter (JSON format)
func (a *RoleController) CreateRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)

	if (len(strings.TrimSpace(role.Code)) == 0) || (len(strings.TrimSpace(role.Name)) == 0) {
		SendBadRequest(c, "Code or Name is empty")
		return
	}

	if err := a.DB.Create(&role).Error; err != nil {
		SendBadRequest(c, "Code is not unique")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"resourceId": role.ID})
}

// UpdateRole change role code or name using id as path parameter
func (a *RoleController) UpdateRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role

	if err := a.DB.Where("id = ?", id).First(&role).Error; err != nil {
		SendNotFound(c, "Role not found")
		return
	}

	c.BindJSON(&role)
	if (len(strings.TrimSpace(role.Code)) == 0) || (len(strings.TrimSpace(role.Name)) == 0) {
		SendBadRequest(c, "Code or Name is empty")
		return
	}
	a.DB.Save(&role)

	c.JSON(http.StatusOK, gin.H{"resourceId": role.ID})
}

// DeleteRole delete role using id as path parameter
func (a *RoleController) DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role

	if err := a.DB.Where("id = ?", id).First(&role).Error; err != nil {
		SendNotFound(c, "Role not found")
		return
	}

	a.DB.Delete(&role)

	c.JSON(http.StatusOK, gin.H{"resourceId": role.ID})
}
