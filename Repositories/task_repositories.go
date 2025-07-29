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

type TaskRepository struct {
	Coll *mongo.Collection
	ctx  context.Context
}

func NewTaskRepository(ctx *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		Coll: ctx,
		ctx:  context.Background(),
	}
}
func (r *TaskRepository) GetAll() ([]domain.Task, error) {
	cur, err := r.Coll.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(r.ctx)
	var tasks []domain.Task
	for cur.Next(r.ctx) {
		var t domain.Task
		if err := cur.Decode(&t); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(id primitive.ObjectID) (*domain.Task, error) {
	var task domain.Task
	if err := r.Coll.FindOne(r.ctx, bson.M{"_id": id}).Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Create(task domain.Task) (*domain.Task, error) {
	task.TaskID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_, err := r.Coll.InsertOne(r.ctx, task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(id primitive.ObjectID, task domain.Task) (*domain.Task, error) {
	task.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"due_date":    task.DueDate,
			"status":      task.Status,
			"updated_at":  task.UpdatedAt,
		},
	}
	res, err := r.Coll.UpdateByID(r.ctx, id, update)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return nil, errors.New("task not found")
	}

	// ✅ Don't reuse `task` name
	updatedTask, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}

func (r *TaskRepository) Delete(id primitive.ObjectID) error {
	res, err := r.Coll.DeleteOne(r.ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
