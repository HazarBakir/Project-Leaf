package main

import (
	"github.com/HazarBakir/Project-Leaf/handlers"
	"github.com/gin-gonic/gin"
)
	
	func main() {
		r := gin.Default()
		r.GET("/", handlers.Homepage)
		r.Run(":8080")
	}
	