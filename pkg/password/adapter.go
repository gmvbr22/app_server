package password

import (
	"golang.org/x/crypto/bcrypt"
)

type Adapter interface {
	GenerateFromPassword(password []byte) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type adapter struct{}

func NewBcryptAdapter() Adapter {
	return &adapter{}
}

func (r *adapter) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, 12)
}

func (r *adapter) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
