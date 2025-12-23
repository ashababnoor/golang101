package main

import (
	"gorm-playground/db"
	"gorm-playground/models"
	"gorm-playground/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	db.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.Run(":8080")
}
