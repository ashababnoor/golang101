package main

import (
	"github.com/gin-gonic/gin"
	"gorm-playground/db"
	"gorm-playground/handlers"
	"gorm-playground/models"
)

func main() {
	db.Connect()

	db.DB.AutoMigrate(
		&models.User{},
		&models.AuditLog{},
	)

	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.POST("/users/tx", handlers.CreateUserWithAudit)

	r.Run(":8080")
}
