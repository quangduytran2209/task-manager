package http

import (
	"net/http"

	"task-manager/internal/config/auth"
	"task-manager/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UU  *usecase.UserUsecase
	JWT auth.JWTService // Add JWT service interface here
}
func NewUserHandler(e *echo.Echo, uu *usecase.UserUsecase, jwt auth.JWTService) {
	h := &UserHandler{UU: uu, JWT: jwt}

	e.POST("/signup", h.SignUp)
	e.POST("/login", h.Login)
}


func (h *UserHandler) SignUp(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error)
	}

	if err := h.UU.SignUp(req.Username, req.Email, req.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "user created")
}

func (h *UserHandler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

	user, err := h.UU.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// todo: call jwt service here
	token, err := h.JWT.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not generate token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
