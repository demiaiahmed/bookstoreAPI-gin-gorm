package controllers

import (
	"net/http"

	"github.com/demiaiahmed/bookstoreAPI-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"Data": books})
}
