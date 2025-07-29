package Usecases

import (
	"errors"

	domain "github.com/surafelbkassa/go-task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCaseInterface interface {
	RegisterUser(name, email, password string) error
	LoginUser(email, password string) (*domain.User, error)
	PromoteUser(userID primitive.ObjectID) (*domain.User, error)
}

// UserUseCase implements user business rules
type userUseCase struct {
	repo   domain.UserRepository
	hasher domain.PasswordHasher
}

// NewUserUseCase constructor
func NewUserUseCase(r domain.UserRepository, h domain.PasswordHasher) *userUseCase {
	return &userUseCase{repo: r, hasher: h}
}

func (uc *userUseCase) RegisterUser(name, email, password string) error {
	// check if user email already exists
	existing, _ := uc.repo.GetByEmail(email)
	if existing != nil {
		return errors.New("email already registered")
	}
	// hash password
	hashedPassword, err := uc.hasher.HashPassword(password)
	if err != nil {
		return err
	}
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "user", // default role
	}

	_, err = uc.repo.Create(*user)
	return err
}

func (uc *userUseCase) LoginUser(email, password string) (*domain.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if user == nil || !uc.hasher.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (uc *userUseCase) PromoteUser(userID primitive.ObjectID) (*domain.User, error) {
	return uc.repo.PromoteUser(userID)
}
