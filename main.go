package main

import (
	"github.com/demiaiahmed/bookstoreAPI-gin-gorm/controllers"
	"github.com/demiaiahmed/bookstoreAPI-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.GetBooks)
	models.ConnectDatabase()
	r.Run()
}
