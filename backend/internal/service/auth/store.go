package auth

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) HashedPassword(password string) (string, error) {
	hPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hPassword), nil
}

func (s *Store) VerifyPassword(hPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hPassword), []byte(password)); err != nil {
		return err
	}

	return nil
}

// func  ()  {

// }
