package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"

    "github.com/AndhikaBhas/miniProject.git/database"
    "github.com/AndhikaBhas/miniProject.git/models"
)

type CreateTaskRequest struct {
    Title       string     `json:"title" binding:"required"`
    Description string     `json:"description"`
    Status      string     `json:"status"`
    Priority    string     `json:"priority"`
    DueDate     *time.Time `json:"due_date"`
    CategoryID  *uuid.UUID `json:"category_id"`
}

type UpdateTaskRequest struct {
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Status      string     `json:"status"`
    Priority    string     `json:"priority"`
    DueDate     *time.Time `json:"due_date"`
    CategoryID  *uuid.UUID `json:"category_id"`
}

func GetTasks(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var tasks []models.Task
	database.DB.
		Preload("Category").
		Where("user_id = ?", userID).
		Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}


func CreateTask(c *gin.Context) {
    userID := c.GetString("user_id")

    var req CreateTaskRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    userUUID, _ := uuid.Parse(userID)
    task := models.Task{
        Title:       req.Title,
        Description: req.Description,
        Status:      req.Status,
        Priority:    req.Priority,
        DueDate:     req.DueDate,
        CategoryID:  req.CategoryID,
        UserID:      userUUID,
    }

    if task.Status == "" {
        task.Status = "todo"
    }
    if task.Priority == "" {
        task.Priority = "medium"
    }

    if err := database.DB.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create task"})
        return
    }

    database.DB.Preload("Category").First(&task, task.ID)

    c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
    userID := c.GetString("user_id")
    taskID := c.Param("id")

    var task models.Task
    if err := database.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }

    var req UpdateTaskRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    if req.Title != "" {
        task.Title = req.Title
    }
    if req.Description != "" {
        task.Description = req.Description
    }
    if req.Status != "" {
        task.Status = req.Status
    }
    if req.Priority != "" {
        task.Priority = req.Priority
    }
    task.DueDate = req.DueDate
    task.CategoryID = req.CategoryID

    if err := database.DB.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update task"})
        return
    }

    database.DB.Preload("Category").First(&task, task.ID)

    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    userID := c.GetString("user_id")
    taskID := c.Param("id")

    result := database.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{})
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete task"})
        return
    }

    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}