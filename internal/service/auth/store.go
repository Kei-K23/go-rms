package auth

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaim struct {
	UserID  int
	Expires int64
	jwt.RegisteredClaims
}
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

func (s *Store) CreateJWT(secret []byte, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		UserID:  userID,
		Expires: int64(time.Second * time.Duration(3600*24*7)),
	})

	t, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return t, nil
}
