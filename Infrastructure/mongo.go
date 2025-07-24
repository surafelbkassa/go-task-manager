package Infrastructure

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient    *mongo.Client
	TaskCollection *mongo.Collection
	UserCollection *mongo.Collection
)

// Connect to MongoDB and assign to MongoClient
func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB")
}

// Initialize task and user collections after MongoClient is connected
func InitMongoUser() {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable not set")
	}

	TaskCollection = MongoClient.Database(dbName).Collection("tasks")
	UserCollection = MongoClient.Database(dbName).Collection("users")

	log.Println("MongoDB collections initialized: tasks, users")
}
