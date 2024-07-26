package main

import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/HazarBakir/Project-Leaf/docs" // Import generated docs
)

var swaggerAuthenticationUsername = "lothlorien"

// @title Project Leaf API
// @version 1.0.0
// @description This is the API for the Project Leaf application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
    router := gin.Default()

    // Public route to generate a token (for demonstration)
    router.POST("/generate-token", getApiToken)


    // Protected routes
    swaggerAuthorized := router.Group("/")
    swaggerAuthorized.Use(authMiddleware)
    {
        swaggerAuthorized.GET("/books", GetBooks)
        swaggerAuthorized.GET("/books/:id", GetBook)
        swaggerAuthorized.POST("/books", CreateBook)
        swaggerAuthorized.PUT("/books/:id", UpdateBook)
        swaggerAuthorized.DELETE("/books/:id", DeleteBook)
    }

    url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    router.Run(":8000")
}
