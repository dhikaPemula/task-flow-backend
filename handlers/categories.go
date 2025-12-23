package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"

    "github.com/AndhikaBhas/miniProject.git/database"
    "github.com/AndhikaBhas/miniProject.git/models"
)

type CreateCategoryRequest struct {
    Name  string `json:"name" binding:"required"`
    Color string `json:"color" binding:"required"`
}

func GetCategories(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var categories []models.Category
	if err := database.DB.
		Where("user_id = ?", userID).
		Find(&categories).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}




func CreateCategory(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user id"})
		return
	}

	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	category := models.Category{
		Name:   req.Name,
		Color:  req.Color,
		UserID: userUUID,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}



func DeleteCategory(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user id"})
		return
	}

	categoryID := c.Param("id")

	result := database.DB.
		Where("id = ? AND user_id = ?", categoryID, userUUID).
		Delete(&models.Category{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
