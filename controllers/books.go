package controllers

import (
	"net/http"

	"github.com/demiaiahmed/bookstoreAPI-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"Data": books})
}
func GetBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Can't Find your Book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func AddBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Can't Add a new Book"})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"Success": "Your Book have been successfully Created"})
}

func DelBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Couldn't Delete Item"})
		return
	}
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"Success": "The item has been deleted successfully"})
}

func EdBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.Title = input.Title
	book.Author = input.Author
	models.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
