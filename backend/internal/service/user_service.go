package service

import (
	"context"
	"errors"
	"fmt" // 👈 ADD THIS
	"taskflow/internal/models"
	"taskflow/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user *models.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	fmt.Println("REGISTER INPUT PASSWORD:", user.Password)
	user.Password = string(hashedPassword)
	fmt.Println("HASHED PASSWORD:", user.Password)

	// Save user
	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println("DB ERROR:", err) // 👈 DEBUG
		return nil, errors.New("invalid credentials")
	}

	fmt.Println("User found:", user.Email)     // 👈 DEBUG
	fmt.Println("Stored hash:", user.Password) // 👈 DEBUG
	fmt.Println("Input password:", password)   // 👈 DEBUG

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password mismatch") // 👈 DEBUG
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
