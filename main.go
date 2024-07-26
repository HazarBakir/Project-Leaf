package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/HazarBakir/Project-Leaf/docs" // Import generated docs
)

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
    router.POST("/generate-token", func(c *gin.Context) {
        username := c.PostForm("username")
        if username == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
            return
        }
        if username == "lothlorien" {
        token := setAPIToken(username)
        c.JSON(http.StatusOK, gin.H{"token": token})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
        }
    })


    // Protected routes
    authorized := router.Group("/")
    authorized.Use(authMiddleware)
    {
        authorized.GET("/books", GetBooks)
        authorized.GET("/books/:id", GetBook)
        authorized.POST("/books", CreateBook)
        authorized.PUT("/books/:id", UpdateBook)
        authorized.DELETE("/books/:id", DeleteBook)
    }

    url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    router.Run(":8000")
}
