package main

import (
	"log"

	"gorm-playground/db"
	"gorm-playground/models"
)

func main() {
	// connect DB
	db.Connect()

	// migrate schema
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	// CREATE
	user := models.User{
		Name:  "Shabab",
		Email: "shabab@example.com",
		Age:   26,
	}

	db.DB.Create(&user)
	log.Println("user created with ID:", user.ID)

	// READ
	var fetched models.User
	db.DB.First(&fetched, "email = ?", "shabab@example.com")
	log.Println("fetched user:", fetched)

	// UPDATE
	db.DB.Model(&fetched).Update("Age", 27)

	// DELETE (soft delete)
	// db.DB.Delete(&fetched)
}
