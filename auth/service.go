package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	// buat payload
	claims := jwt.MapClaims{}
	claims["user_id"] = userId

	// claim dan sign jwt
	JWT_SECRET_KEY := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
