package handlers

import (
	"net/http"

	"gorm-playground/db"
	"gorm-playground/dto"
	"gorm-playground/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserWithAudit(c *gin.Context) {
	var req dto.CreateUserRequest

	// 1. Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var createdUser models.User

	// 2. BEGIN TRANSACTION
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// create user
		user := models.User{
			Name:  req.Name,
			Email: req.Email,
			Age:   req.Age,
		}

		if err := tx.Create(&user).Error; err != nil {
			return err // ROLLBACK
		}

		// create audit log
		audit := models.AuditLog{
			Entity:   "user",
			EntityID: user.ID,
			Action:   "create",
		}

		if err := tx.Create(&audit).Error; err != nil {
			return err // ROLLBACK
		}

		createdUser = user
		return nil // COMMIT
	})

	// 3. Handle transaction result
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "transaction failed: " + err.Error(),
		})
		return
	}

	// 4. Respond with DTO
	resp := dto.CreateUserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
		Age:   createdUser.Age,
	}

	c.JSON(http.StatusCreated, resp)
}
