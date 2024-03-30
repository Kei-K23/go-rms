package payment

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

func (s *Store) CreatePayment(payment types.CreatePayment) error {
	stmt, err := s.db.Prepare("INSERT INTO payments (order_id, amount, payment_type) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(payment.OrderID, payment.Amount, payment.PaymentType)
	if err != nil {
		return err
	}

	return nil
}
