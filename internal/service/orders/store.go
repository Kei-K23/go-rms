package orders

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-rms/backend/internal/types"
)

type Store struct {
	db      *sql.DB
	oiStore types.OrderItemStore
}

func NewStore(db *sql.DB, oiStore types.OrderItemStore) *Store {
	return &Store{db: db, oiStore: oiStore}
}

func (s *Store) GetOrderByID(id int) (*types.Order, error) {
	var order types.Order

	stmt, err := s.db.Prepare("SELECT * FROM orders WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = stmt.QueryRow(id).Scan(&order.ID, &order.TableNumber, &order.TotalPrice, &order.TotalQuantity, &order.OrderTime, &order.RestaurantId, &order.OrderStatus, &order.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no order found")
	}
	return &order, nil
}

func (s *Store) CreateOrder(order types.CreateOrder, orderItems []types.CreateOrderItem) (*types.Order, error) {
	stmt, err := s.db.Prepare("INSERT INTO orders (table_number, total_price, total_quantity, restaurant_id, order_status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	defer stmt.Close()

	result, err := stmt.Exec(order.TableNumber, order.TotalPrice, order.TotalQuantity, order.RestaurantId, order.OrderStatus)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	oID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	createdOrder, err := s.GetOrderByID(int(oID))
	if err != nil {
		return nil, err
	}
	// create order item here
	for _, oi := range orderItems {
		oi.OrderID = createdOrder.ID
		_, err := s.oiStore.CreateOrderItem(oi)
		if err != nil {
			return nil, fmt.Errorf("internal server error")
		}
	}

	return createdOrder, nil
}

func (s *Store) DeleteOrder(oID, rID int) (*types.HTTPGeneralRes, error) {
	_, err := s.GetOrderByID(oID)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare("DELETE FROM orders WHERE id = ? AND restaurant_id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(oID, rID)
	if err != nil {
		return nil, err
	}

	return &types.HTTPGeneralRes{Success: true, Message: "Deleted order with ID: " + fmt.Sprintf("%d", oID)}, nil
}

// change order status
func (s *Store) ChangeOrderStatus(id int) (*types.HTTPGeneralRes, error) {
	order, err := s.GetOrderByID(id)
	if err != nil {
		return nil, fmt.Errorf("no order item found")
	}

	var orderS types.OrderStatus

	stmt, err := s.db.Prepare("UPDATE orders SET order_status = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	if order.OrderStatus == "pending" {
		orderS = "completed"
	} else {
		orderS = "pending"
	}

	_, err = stmt.Exec(orderS, id)
	if err != nil {
		return nil, fmt.Errorf("no order found")
	}
	return &types.HTTPGeneralRes{Message: "order status updated successfully", Success: true}, nil
}
