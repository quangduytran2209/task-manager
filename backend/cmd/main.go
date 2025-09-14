package main

import (
	"task-manager/internal/config/auth"
	"task-manager/internal/infrastructure/db"
	"task-manager/internal/infrastructure/http"
	"task-manager/internal/usecase"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"  
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB → Repository → Usecase → Handler → Echo server
	// 1. connect to DB
	// Connect to PostgreSQL database
	dsn := "host=localhost user=taskuser password=taskpass dbname=taskbd port=5432 sslmode=disable"
	gormDB, err := gorm.Open((postgres.Open(dsn)), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 2. migrate tables
	// run auto migration to create/update tables based on models
	gormDB.AutoMigrate(&db.TaskModel{})

	// 3. repository and usecase
	// create repository and usecase instances
	taskRepo := db.NewPostgresTaskRepository(gormDB)
	userRepo := db.NewPostgresUserRepository(gormDB)

	// 4. usecase
	// inject repository into usecase
	taskUC := usecase.NewTaskUsecase(taskRepo)
	userUC := usecase.NewUserUsecase(userRepo)

	// create jwt service
	jwtService := auth.NewJWTService("secret-key", "task-manager")

	// 5. echo server
	e := echo.New()
	
	// 6. handlers
	r := e.Group("/api")
	r.Use(echojwt.WithConfig(echojwt.Config{
	SigningKey:  []byte("secret-key"),
	TokenLookup: "header:Authorization",
}))
	http.NewTaskHandler(e, taskUC)
	http.NewUserHandler(e, userUC, jwtService)

	// 7. start server
	e.Logger.Fatal(e.Start(":8080"))
}
