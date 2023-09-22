package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	structs "rest"
	"rest/pkg/repository"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "ghfghgd5345345fdds"
	signingKey = "grgdfgfd$gfd654#fgdsdf$fd"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user structs.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	fmt.Println(user, "Into serice")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	},
	)
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
