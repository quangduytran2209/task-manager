package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"task-manager/internal/database"
	"task-manager/internal/models"
)

// use ENV, sercrrtkey
var jwtSecret = []byte("sercretkey")

// Register
func Register(c echo.Context) error {
	u := new(models.User)

	//get request
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invaild input"})
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "fail to hash pasword"})
	}
	u.Password = string(hash)

	//check mail/username
	if err := database.DB.Create(u).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "user already exists"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "user name registered"})
}

// Login
func Login(c echo.Context) error {

	//temp struct get data
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//check login
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invaild input"})
	}

	//find user by email
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invaild email or password"})
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invaild email or password"})
	}

	//create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error":"cant sign token"})
	}
	return c.JSON(http.StatusOK, echo.Map{"token":t})
}
