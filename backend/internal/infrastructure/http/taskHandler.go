package http

import (
	"net/http"
	"time"

	"task-manager/internal/domain"
	"task-manager/internal/usecase"

	"github.com/labstack/echo/v4"
)

// TaskHandler calls usecase methods and handles HTTP requests/responses meaning it acts as a bridge between the HTTP layer and the usecase layer
type TaskHandler struct {
	UC *usecase.TaskUsecase
}

// NewTaskHandler registers task-related routes to the Echo instance meaning it sets up the endpoints for task operations
func NewTaskHandler(e *echo.Echo, uc *usecase.TaskUsecase) {
	h := &TaskHandler{UC: uc}
	g := e.Group("/api")
	g.GET("/tasks", h.ListByUser)
	g.POST("/tasks", h.CreateTask)
	// g.PUT("/tasks/:id", h.UpdateTask)
	// g.DELETE("/tasks/:id", h.DeleteTask)
}

// ListByUser handles GET /tasks and returns tasks for a user
func (h *TaskHandler) ListByUser(c echo.Context) error {
	// (Later) extract UserID from JWT token/session
	userID := uint(1) // hardcoded for now

	tasks, err := h.UC.ListTasksByUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.(error).Error()})
	}
	return c.JSON(http.StatusOK, tasks)
}

// create endpoint to create a new task meaning handle POST /tasks
func (h *TaskHandler) CreateTask(c echo.Context) error {
	var dto struct { // Data Transfer Object for creating a task
		Title       string     `json:"title" validate:"required"`
		Description string     `json:"description"`
		Status      string     `json:"status"`
		Deadline    *time.Time `json:"deadline"`
		UserID      uint       `json:"user_id" validate:"required"`
	}
	// parese request body meaning bind json to dto
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	// 2. (Later) extract UserID from JWT token/session meaning get user ID from auth context
	userID := uint(1) // later replace with JWT meaning hardcoded for now

	// 3. create domain task entity meaning map dto to domain entity meaning prepare data for usecase
	task := &domain.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		Deadline:    dto.Deadline,
		UserID:      userID,
	}

	// 4. call usecase to create task meaning invoke business logic
	if err := h.UC.CreateTask(task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// send response meaning return created task with 201 status
	return c.JSON(http.StatusCreated, task)
}
