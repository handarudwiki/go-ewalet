package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/joho/godotenv"
)

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(user domain.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	key := []byte(os.Getenv("JWT_SECRET_KEY"))

	claim := jwt.MapClaims{}
	claim["user_id"] = user.ID
	claim["name"] = user.Name
	claim["username"] = user.Username
	claim["phone"] = user.Phone

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return string(signedToken), nil
}
