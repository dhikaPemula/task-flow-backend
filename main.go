package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/AndhikaBhas/miniProject.git/config"
	"github.com/AndhikaBhas/miniProject.git/database"
	"github.com/AndhikaBhas/miniProject.git/handlers"
	"github.com/AndhikaBhas/miniProject.git/middleware"
	"github.com/AndhikaBhas/miniProject.git/utils"
)

func main() {
	godotenv.Load()

	cfg := config.LoadConfig()

	utils.InitJWT(cfg.JWTSecret)
	database.Connect(cfg.DatabaseURL)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			//categories
			protected.GET("/categories", handlers.GetCategories)
			protected.POST("/categories", handlers.CreateCategory)
			protected.DELETE("/categories/:id", handlers.DeleteCategory)

			//tasks
			protected.GET("/tasks", handlers.GetTasks)
			protected.POST("/tasks", handlers.CreateTask)
			protected.PUT("/tasks/:id", handlers.UpdateTask)
			protected.DELETE("/tasks/:id", handlers.DeleteTask)
		}
	}

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
