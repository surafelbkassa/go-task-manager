package Usecases

import (
	"context"
	"errors"

	domain "github.com/surafelbkassa/go-task-manager/Domain"
	"github.com/surafelbkassa/go-task-manager/Infrastructure"
	"github.com/surafelbkassa/go-task-manager/Repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCaseInterface interface {
	RegisterUser(domain.User) (*domain.User, error)
	LoginUser(email, password string) (*domain.User, error)
	PromoteUser(userID string) error
}

// UserUseCase implements user business rules
type UserUseCase struct {
	repo *Repositories.UserRepository
}

// NewUserUseCase constructor
func NewUserUseCase(r *Repositories.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) RegisterUser(input domain.User) (*domain.User, error) {
	// check if user email already exists
	existing, _ := u.repo.GetByEmail(input.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}
	// hash password
	hash, err := Infrastructure.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	input.Password = hash
	// assign role: first user is admin, rest are users
	count, _ := u.repo.Coll.CountDocuments(context.Background(), bson.D{})
	if count == 0 {
		input.Role = "admin"
	} else {
		input.Role = "user"
	}
	return u.repo.Register(input)
}

func (u *UserUseCase) LoginUser(email, password string) (*domain.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	err = Infrastructure.CheckPassword(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (u *UserUseCase) PromoteUser(userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid ID format")
	}
	return u.repo.Promote(objID)
}
