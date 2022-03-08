package main

import (
	"github.com/demiaiahmed/bookstoreAPI-gin-gorm/controllers"
	"github.com/demiaiahmed/bookstoreAPI-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.POST("/add", controllers.AddBook)
	r.DELETE("/books/:id", controllers.DelBook)
	r.PATCH("/books/:id", controllers.EdBook)

	r.Run()
}
