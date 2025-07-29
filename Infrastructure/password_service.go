package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}
func (s *PasswordService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (ps *PasswordService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
