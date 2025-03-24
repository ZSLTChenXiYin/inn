package utils

import "golang.org/x/crypto/bcrypt"

const (
	PASSWORD_HASH_COST = 12
)

func GenerateHashPassword(password string) (string, error) {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), PASSWORD_HASH_COST)
	return string(password_hash), err
}

func ValidatePassword(password string, password_hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
}
