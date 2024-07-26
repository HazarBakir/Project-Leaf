package main

import (
    "crypto/rand"
    "encoding/hex"
    "net/http"
    "github.com/gin-gonic/gin"
)

type MessageResponse struct {
    Message string `json:"message"`
}

// Single API token (for demonstration purposes)
var apiToken string

// Function to generate a new API token
func generateAPIToken() string {
    token := make([]byte, 16)
    _, err := rand.Read(token)
    if err != nil {
        panic(err)
    }
    return hex.EncodeToString(token)
}

// Function to set a new token for the given username
func setAPIToken(username string) string {
    apiToken = generateAPIToken()
    return apiToken
}

// Middleware to authenticate API requests
func authMiddleware(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusUnauthorized, MessageResponse{Message: "API token required"})
        c.Abort()
        return
    }

    if token != apiToken {
        c.JSON(http.StatusUnauthorized, MessageResponse{Message: "Invalid API token"})
        c.Abort()
        return
    }

    c.Next()
}
