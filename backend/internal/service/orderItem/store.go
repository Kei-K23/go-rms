package orderitem

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

func (s *Store) GetOrderItemByID(id int) (*types.OrderItem, error) {
	var orderItem types.OrderItem

	stmt, err := s.db.Prepare("SELECT * FROM order_items WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = stmt.QueryRow(id).Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.MenuID, &orderItem.Quantity, &orderItem.Price, &orderItem.OrderStatus, &orderItem.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no order item found")
	}
	return &orderItem, nil
}

func (s *Store) CreateOrderItem(orderItem types.CreateOrderItem) (*types.OrderItem, error) {
	stmt, err := s.db.Prepare("INSERT INTO order_items (order_id, menu_id, quantity, price, order_status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(orderItem.OrderID, orderItem.MenuID, orderItem.Quantity, orderItem.Price, orderItem.OrderStatus)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error when creating order item: %v", err)
	}

	oI, err := s.GetOrderItemByID(int(id))
	if err != nil {
		return nil, err
	}

	return oI, nil
}
