package main

import (
	"github.com/Polluxe/GolangRestApi/controllers"
	"github.com/Polluxe/GolangRestApi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	duty := gin.Default()
	models.ConnectDatabase()

	duty.GET("/api/items", itemcontroller.Index)
	duty.GET("/api/item/:id", itemcontroller.Show)
	duty.POST("/api/item", itemcontroller.Create)
	duty.PUT("/api/item/:id", itemcontroller.Update)
	duty.DELETE("/api/item", itemcontroller.Delete)
}
