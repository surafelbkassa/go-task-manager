package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/surafelbkassa/go-task-manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func InitMongo() {
	// var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, error := mongo.Connect(context.TODO(), clientOptions)

	if error != nil {
		log.Fatal(error)
	}
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
	}
	TaskCollection = client.Database("task_manager").Collection("tasks")
	fmt.Println("Connected to MongoDB!")

}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := TaskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}
	var task models.Task
	filter := bson.M{"_id": objID}
	err = TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdatedTask(id string, updatedTask models.Task) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}
	update := bson.M{}
	if updatedTask.Title != "" {
		update["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		update["description"] = updatedTask.Description
	}
	if !updatedTask.DueDate.IsZero() {
		update["due_date"] = updatedTask.DueDate
	}
	if updatedTask.Status != "" {
		update["status"] = updatedTask.Status
	}
	if len(update) == 0 {
		return nil, errors.New("no fields to update")
	}
	filter := bson.M{"_id": objID}
	_, err = TaskCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return GetTaskById(id)
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}
	filter := bson.M{"_id": objID}
	result, err := TaskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
func CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_, err := TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	return &task, nil

}
