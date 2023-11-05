package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/skripov-ds-ai/social_network/internal/database/user"
	"github.com/skripov-ds-ai/social_network/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Get(ctx context.Context, userID string) (models.User, error)
	Login(ctx context.Context, userID, password string) (string, error)
	Register(ctx context.Context, params models.CreateUserParams) (string, error)
}

type Service struct {
	repository user.Repository
}

func NewUserService(repository user.Repository) UserService {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Get(ctx context.Context, userID string) (models.User, error) {
	return s.repository.Get(ctx, userID)
}

func (s *Service) Login(ctx context.Context, userID, password string) (string, error) {
	passwordHash, err := hashPassword(password)
	if err != nil {
		return "", err
	}

	u, err := s.repository.Get(ctx, userID)
	if err != nil {
		return "", err
	}

	// TODO
	if passwordHash != u.PasswordHash {
		return "", errors.New("wrong password")
	}

	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *Service) Register(ctx context.Context, params models.CreateUserParams) (string, error) {
	passwordHash, err := hashPassword(params.Password)
	if err != nil {
		return "", err
	}
	return s.repository.Create(ctx, models.User{
		PasswordHash: passwordHash,
		FirstName:    params.FirstName,
		SecondName:   params.SecondName,
		Birthdate:    params.Birthdate,
		Biography:    params.Biography,
		City:         params.City,
	})
}
