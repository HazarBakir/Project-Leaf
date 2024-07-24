package main

import (
	"github.com/HazarBakir/Project-Leaf/handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"github.com/HazarBakir/Project-Leaf/cmd/docs"
)
	
// @title Gin CRUD API
// @version 1.0
// @description This is a sample CRUD API using Gin.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/", handlers.Homepage)
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBook)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		
	router.Run(":8080")
}
	