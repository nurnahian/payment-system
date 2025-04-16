package user

import (
	"database/sql"
	"errors"
	"fmt"
	"payment-system/internal/auth"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
	GetAllUsers() ([]User, error)
}

type userService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &userService{repo: r}
}

func (s *userService) Register(username, password string) error {
	_, err := s.repo.FindByUsername(username)
	if err == nil {
		// User already exists
		return fmt.Errorf("user already exists")
	} else if !errors.Is(err, sql.ErrNoRows) {
		// Something went wrong while checking
		return err
	}

	// Hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create new user
	return s.repo.Create(username, string(hashed))
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return auth.GenerateJWT(user.Username)
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}
