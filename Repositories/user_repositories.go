package Repositories

import (
	"context"
	"time"

	domain "github.com/surafelbkassa/go-task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Coll *mongo.Collection
}

func NewUserRepository(c *mongo.Collection) *UserRepository {
	return &UserRepository{Coll: c}
}

func (r *UserRepository) Register(user domain.User) (*domain.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := r.Coll.InsertOne(ctx, user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user domain.User
	if err := r.Coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Promote(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Coll.UpdateByID(ctx, id, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}
