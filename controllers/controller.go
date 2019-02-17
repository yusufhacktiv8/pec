package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
}

func SendBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
}
