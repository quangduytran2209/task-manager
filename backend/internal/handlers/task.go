package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"task-manager/internal/database"
	"task-manager/internal/models"
)

// help to extract user_id from JWT
func getUserID(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)    //get token from context
	claims := user.Claims.(jwt.MapClaims) // parse claim (date in token)

	return uint(claims["user_id"].(float64)) // typecast and return user_id
}

// getTask.
func GetTasks(c echo.Context) error {
	userID := getUserID(c) // get user id from token

	var tasks []models.Task
	database.DB.Where("user_id =?", userID).Find(&tasks) // findding all task by user id

	return c.JSON(http.StatusOK, tasks)
}

// postTask
func CreateTasks(c echo.Context) error {
	userID := getUserID(c)

	var task models.Task
	if err := c.Bind(&task); err != nil { // parse body JSON -> struct Task
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invaild input"})
	}

	task.UserID = userID // assign user for task
	if task.Status == "" {
		task.Status = "todo"
	}

	if task.Deadline.IsZero() {
		task.Deadline = time.Now().Add(24 * time.Hour)
	}

	database.DB.Create(&task) // save db
	return c.JSON(http.StatusOK, task)

}

// putTask
func UpdateTasks(c echo.Context) error {
	userID := getUserID(c)
	id := c.Param("id") //  get param from URL

	var task models.Task
	if err := database.DB.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "task not found"})
	}

	if err := c.Bind(&task); err != nil { //update infor by body
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invaild input"})
	}

	database.DB.Save(&task) // save data
	return c.JSON(http.StatusOK, task)
}

// deateleTask
func DeateleTasks(c echo.Context) error {
	userID := getUserID(c)
	id := c.Param("id")

	var task models.Task
	if err := database.DB.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "task not found"})
	}

	database.DB.Delete(&task) // save data
	return c.JSON(http.StatusOK, task)
}
