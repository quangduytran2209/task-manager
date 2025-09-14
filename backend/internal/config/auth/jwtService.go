package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	isUser    string
}

func NewJWTService(secretKey, isUser string) JWTService {
	return &jwtService{secretKey: secretKey, isUser: isUser}
}

func (j *jwtService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["isUser"] = j.isUser
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // token expires in 72 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}
