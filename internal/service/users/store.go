package users

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-rms/backend/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.RegisterUser) (*types.RegisterUser, error) {
	stmt, err := s.db.Prepare("INSERT INTO users (name, email, password, phone, address, access_key) VALUES (? , ? , ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.Phone, user.Address, user.AccessKey)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) GetUserByEmail(user types.LoginUser) (*types.User, error) {
	var u types.User

	stmt, err := s.db.Prepare("SELECT * FROM users WHERE email =?")
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	defer stmt.Close()

	err = stmt.QueryRow(user.Email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Address, &u.Phone, &u.AccessKey, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("user could not found")
	}

	return &u, nil
}

func (s *Store) GetUserById(uID int) (*types.User, error) {
	var u types.User

	stmt, err := s.db.Prepare("SELECT * FROM users WHERE id =?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(uID).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Address, &u.Phone, &u.AccessKey, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Store) UpdateUser(user types.UpdateUser, uID int) (*types.User, error) {
	_, err := s.GetUserById(uID)
	if err != nil {
		return nil, err
	}

	var u *types.User

	stmt, err := s.db.Prepare("UPDATE users SET name = ?, address = ?, phone = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Address, user.Phone, uID)
	if err != nil {
		return nil, err
	}

	u, err = s.GetUserById(uID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) DeleteUser(uID int) (*types.HTTPGeneralRes, error) {
	_, err := s.GetUserById(uID)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uID)
	if err != nil {
		return nil, err
	}

	return &types.HTTPGeneralRes{
		Success: true,
		Message: fmt.Sprintf("User with ID : %v deleted successfully", uID),
	}, nil
}
