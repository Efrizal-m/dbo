package services

import (
	"dbo/models"
	"dbo/repositories"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user models.User) (models.User, error)
	Login(username, password string) (string, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
