package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
