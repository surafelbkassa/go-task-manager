package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ???
type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id primitive.ObjectID) (*Task, error)
	Create(Task) (*Task, error)
	Update(id primitive.ObjectID, task Task) (*Task, error)
	Delete(id primitive.ObjectID) error
}

type UserRepository interface {
	Create(user User) (*User, error)
	GetByID(id primitive.ObjectID) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
	//what is the return []*User and *User difference?
	// Update(id primitive.ObjectID, user User) (*User, error)
	// Delete(id primitive.ObjectID) error
	// Login(email, password primitive.ObjectID) (string, error)
	PromoteUser(id primitive.ObjectID) (*User, error)
}

// ???
type Task struct {
	TaskID      primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     time.Time          `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type User struct {
	UserID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Role      string             `json:"role" bson:"role"` // e.g., "admin", "user"
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type JWTService interface {
	GenerateToken(userID primitive.ObjectID, role string) (string, error)
	ValidateToken(token string) (*primitive.ObjectID, string, error)
}
