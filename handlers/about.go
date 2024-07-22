package handlers

import (
	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "About",
	})
}