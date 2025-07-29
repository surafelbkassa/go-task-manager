package Repositories

import (
	"context"
	"errors"
	"time"

	domain "github.com/surafelbkassa/go-task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// compile‑time check that UserRepository implements domain.UserRepository
var _ domain.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	Coll *mongo.Collection
	ctx  context.Context
}

func NewUserRepository(c *mongo.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{
		Coll: c,
		ctx:  ctx,
	}
}

func (r *UserRepository) Create(user domain.User) (*domain.User, error) {
	user.UserID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	_, err := r.Coll.InsertOne(r.ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	err := r.Coll.FindOne(r.ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.Coll.FindOne(r.ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll() ([]*domain.User, error) {
	cur, err := r.Coll.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(r.ctx)

	var users []*domain.User
	for cur.Next(r.ctx) {
		var u domain.User
		if err := cur.Decode(&u); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *UserRepository) PromoteUser(id primitive.ObjectID) (*domain.User, error) {
	res, err := r.Coll.UpdateOne(
		r.ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"role": "admin"}},
	)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return nil, errors.New("user not found")
	}
	return r.GetByID(id)
}
