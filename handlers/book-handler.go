package handlers

import(
	"github.com/HazarBakir/Project-Leaf/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context){
	c.JSON(http.StatusOK, models.Books)
}


// GetBook godoc
// @Summary Get a book by ID
// @Description Get details of a book by its ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} gin.H{"message": "book not found"}
// @Router /books/{id} [get]
func GetBook(c *gin.Context){
	id := c.Param("id")
	for _, book := range models.Books{
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book"
// @Success 201 {object} models.Book
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Router /books [post]
func CreateBook(c *gin.Context){
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.Books = append(models.Books, book)
	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update the details of an existing book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 404 {object} gin.H{"message": "book not found"}
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range models.Books {
		if book.ID == id {
			models.Books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book from the collection by its ID
// @Tags books
// @Param id path string true "Book ID"
// @Success 200 {object} gin.H{"message": "book deleted"}
// @Failure 404 {object} gin.H{"message": "book not found"}
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range models.Books {
		if book.ID == id {
			models.Books = append(models.Books[:i], models.Books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}