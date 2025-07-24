package Repositories

import (
	"context"
	"time"

	domain "github.com/surafelbkassa/go-task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	Coll *mongo.Collection
}

func NewTaskRepository(c *mongo.Collection) *TaskRepository {
	return &TaskRepository{Coll: c}
}

func (r *TaskRepository) GetAll() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := r.Coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var tasks []domain.Task
	if err := cur.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(id primitive.ObjectID) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var task domain.Task
	if err := r.Coll.FindOne(ctx, bson.M{"_id": id}).Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Create(task domain.Task) (*domain.Task, error) {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := r.Coll.InsertOne(ctx, task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(id primitive.ObjectID, task domain.Task) (*domain.Task, error) {
	task.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{"$set": task}
	if _, err := r.Coll.UpdateByID(ctx, id, update); err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *TaskRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := r.Coll.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}
