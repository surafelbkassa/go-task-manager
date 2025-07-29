package Infrastructure

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTService struct {
	secretKey string
	expiry    time.Duration
}
type JWTServiceInterface interface {
	GenerateToken(userID primitive.ObjectID, role string) (string, error)
	ValidateToken(tokenStr string) (*primitive.ObjectID, string, error)
}

func NewJWTService(secret string, duration time.Duration) JWTServiceInterface {
	return &JWTService{
		secretKey: secret,
		expiry:    duration,
	}
}

func (j *JWTService) GenerateToken(userID primitive.ObjectID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.Hex(),
		"role":    role,
		"exp":     time.Now().Add(j.expiry).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(tokenStr string) (*primitive.ObjectID, string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, "", errors.New("invalid claims")
	}

	idStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, "", errors.New("user ID missing in token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, "", errors.New("role missing in token")
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, "", errors.New("invalid user ID in token")
	}

	return &id, role, nil
}
