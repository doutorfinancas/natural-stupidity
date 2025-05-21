package service

import (
	"errors"
	"os"
	"time"

	"github.com/doutorfinancas/natural-stupidity/repository"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// AuthService handles user authentication and token generation.
type AuthService interface {
	Authenticate(email, password string) (*model.User, error)
	GenerateToken(user *model.User) (string, error)
}

type authService struct {
	repo      repository.UserRepository
	jwtSecret string
}

// NewAuthService creates a new AuthService.
func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo, jwtSecret: os.Getenv("JWT_SECRET")}
}

// Authenticate checks credentials and returns the user if valid.
func (a *authService) Authenticate(email, password string) (*model.User, error) {
	user, err := a.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

// GenerateToken issues a JWT for the authenticated user.
func (a *authService) GenerateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}
