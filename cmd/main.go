package main

import (
	"github.com/HazarBakir/Project-Leaf/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/", handlers.Homepage)
	r.Run(":8080")
}