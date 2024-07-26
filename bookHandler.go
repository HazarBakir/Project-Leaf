package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Produce json
// @Success 200 {array} Book
// @Security ApiKeyAuth
// @Router /books [get]
func GetBooks(c *gin.Context) {
    c.JSON(http.StatusOK, Books)
}

// GetBook godoc
// @Summary Get a book by ID
// @Description Get details of a book by its ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} MessageResponse
// @Security ApiKeyAuth
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
    id := c.Param("id")

    for _, book := range Books {
        if book.Id == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }
    c.JSON(http.StatusNotFound, MessageResponse{Message: "book not found"})
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body Book true "Book"
// @Success 201 {object} Book
// @Failure 400 {object} MessageResponse
// @Security ApiKeyAuth
// @Router /books [post]
func CreateBook(c *gin.Context) {
    var newBook Book

    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, MessageResponse{Message: err.Error()})
        return
    }

    Books = append(Books, newBook)
    c.JSON(http.StatusCreated, newBook)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update the details of an existing book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body Book true "Book"
// @Success 200 {object} Book
// @Failure 400 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Security ApiKeyAuth
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book

    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, MessageResponse{Message: err.Error()})
        return
    }

    for i, book := range Books {
        if book.Id == id {
            Books[i] = updatedBook
            c.JSON(http.StatusOK, updatedBook)
            return
        }
    }
    c.JSON(http.StatusNotFound, MessageResponse{Message: "book not found"})
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book from the collection by its ID
// @Tags books
// @Param id path string true "Book ID"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Security ApiKeyAuth
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
    id := c.Param("id")

    for i, book := range Books {
        if book.Id == id {
            Books = append(Books[:i], Books[i+1:]...)
            c.JSON(http.StatusOK, MessageResponse{Message: "book deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, MessageResponse{Message: "book not found"})
}
