package users

import (
	"database/sql"

	"github.com/Kei-K23/go-rms/backend/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.RegisterUser) (*types.RegisterUser, error) {

	stmt, err := s.db.Prepare("INSERT INTO users (name, email, password, phone, address) VALUES (? , ? , ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.Phone, user.Address)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
