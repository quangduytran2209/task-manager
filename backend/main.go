package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	

	"task-manager/internal/database"
	"task-manager/internal/handlers"
	"task-manager/internal/middleware"
	"task-manager/internal/models"
)

func main() {
	e := echo.New() //create new variable for echo
	//connect to DB
	database.Connect()
	//migrate schema
	database.DB.AutoMigrate(&models.User{})

	//Auth router
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	//Protected routes
	r := e.Group("/api")
	r.Use(echojwt.WithConfig(middleware.JWTMiddleware()))

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTasks)
	r.PUT("/tasks/:id", handlers.DeateleTasks)
	r.DELETE("/tasks/:id", handlers.DeateleTasks)
	
	

	e.Logger.Fatal(e.Start(":8080")) //start server local with port 8080, Note:e.logger.Fatal meaning announ when server have problem
}
