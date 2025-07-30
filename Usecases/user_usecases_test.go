package Usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	Domain "github.com/surafelbkassa/go-task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// --- Mock UserRepository ---
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetByID(id primitive.ObjectID) (*Domain.User, error) {
	args := m.Called(id)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*Domain.User), args.Error(1)
}

func (m *MockUserRepo) Create(u Domain.User) (*Domain.User, error) {
	args := m.Called(u)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*Domain.User), args.Error(1)
}

func (m *MockUserRepo) GetByEmail(email string) (*Domain.User, error) {
	args := m.Called(email)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*Domain.User), args.Error(1)
}

func (m *MockUserRepo) PromoteUser(id primitive.ObjectID) (*Domain.User, error) {
	args := m.Called(id)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*Domain.User), args.Error(1)
}

func (m *MockUserRepo) GetAll() ([]*Domain.User, error) {
	args := m.Called()
	users := args.Get(0)
	if users == nil {
		return nil, args.Error(1)
	}
	return users.([]*Domain.User), args.Error(1)
}

// --- Mock PasswordHasher ---
type MockHasher struct {
	mock.Mock
}

func (m *MockHasher) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockHasher) CheckPasswordHash(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}

// --- Tests ---

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	mockRepo.On("GetByEmail", "a@b.com").Return(nil, nil)
	mockHash.On("HashPassword", "pw").Return("hashed", nil)
	mockRepo.On("Create", Domain.User{
		Name:     "Alice",
		Email:    "a@b.com",
		Password: "hashed",
		Role:     "user",
	}).Return(&Domain.User{Email: "a@b.com"}, nil)

	err := uc.RegisterUser("Alice", "a@b.com", "pw")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
	mockHash.AssertExpectations(t)
}

func TestRegisterUser_ExistingEmail(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	mockRepo.On("GetByEmail", "a@b.com").Return(&Domain.User{}, nil)

	err := uc.RegisterUser("Alice", "a@b.com", "pw")
	assert.EqualError(t, err, "email already registered")
}

func TestLoginUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	stored := &Domain.User{Password: "hash"}
	mockRepo.On("GetByEmail", "e@x.com").Return(stored, nil)
	mockHash.On("CheckPasswordHash", "pw", "hash").Return(true)

	user, err := uc.LoginUser("e@x.com", "pw")
	assert.NoError(t, err)
	assert.Equal(t, stored, user)
}

func TestLoginUser_Fail(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	mockRepo.On("GetByEmail", "e@x.com").Return(nil, errors.New("not found"))

	user, err := uc.LoginUser("e@x.com", "pw")
	assert.Nil(t, user)
	assert.EqualError(t, err, "invalid credentials")
}

func TestPromoteUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	id := primitive.NewObjectID()
	expected := &Domain.User{Email: "z@z.com"}

	mockRepo.On("PromoteUser", id).Return(expected, nil)

	user, err := uc.PromoteUser(id)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}

func TestPromoteUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	id := primitive.NewObjectID()
	mockRepo.On("PromoteUser", id).Return(nil, errors.New("oops"))

	_, err := uc.PromoteUser(id)
	assert.EqualError(t, err, "oops")
}
func TestRegisterUser_HashError(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	mockRepo.On("GetByEmail", "a@b.com").Return(nil, nil)
	mockHash.On("HashPassword", "pw").Return("", errors.New("hash failed"))

	err := uc.RegisterUser("Alice", "a@b.com", "pw")
	assert.EqualError(t, err, "hash failed")

	mockRepo.AssertExpectations(t)
	mockHash.AssertExpectations(t)
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	stored := &Domain.User{Password: "hash"}
	mockRepo.On("GetByEmail", "e@x.com").Return(stored, nil)
	mockHash.On("CheckPasswordHash", "pw", "hash").Return(false)

	user, err := uc.LoginUser("e@x.com", "pw")
	assert.Nil(t, user)
	assert.EqualError(t, err, "invalid credentials")
}

func TestLoginUser_UserNilButNoError(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHash := new(MockHasher)
	uc := NewUserUseCase(mockRepo, mockHash)

	mockRepo.On("GetByEmail", "e@x.com").Return(nil, nil)

	user, err := uc.LoginUser("e@x.com", "pw")
	assert.Nil(t, user)
	assert.EqualError(t, err, "invalid credentials")
}
