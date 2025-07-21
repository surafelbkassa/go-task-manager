package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/surafelbkassa/go-task-manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var Collection *mongo.Collection

func InitMongoUser() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	error := client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
	}
	Collection = client.Database("task_manager").Collection("users")
	fmt.Println("Connected to MongoDB for User Service!")
}

func RegisterUser(name, email, password string) (*models.User, error) {
	// Check if email exists
	var existing models.User
	err := Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&existing)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Count users in the DB to make the first one admin
	count, err := Collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	role := "user"
	if count == 0 {
		role = "admin"
	}

	user := models.User{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: time.Now(),
	}

	_, err = Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ✅ LoginUser authenticates a user and returns a JWT
func LoginUser(email, password string) (string, error) {
	var user models.User
	err := Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	return GenerateJWT(&user)
}
func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
func PromoteUserToAdmin(userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err = Collection.UpdateOne(context.TODO(), filter, update)
	return err
}
