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

func (s *Store) GetOrderItemByOrderID(orderID int) (*[]types.OrderItem, error) {
	var orderItems []types.OrderItem

	stmt, err := s.db.Prepare("SELECT * FROM order_items WHERE order_id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	result, err := stmt.Query(orderID)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no order item found")
	}

	if result.Next() {
		var orderItem types.OrderItem
		result.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.MenuID, &orderItem.Quantity, &orderItem.Price, &orderItem.OrderStatus, &orderItem.CreatedAt)
		orderItems = append(orderItems, orderItem)
	}
	return &orderItems, nil
}

func (s *Store) UpdateOrderItem(oiUpdate types.UpdateOrderItem, id int) (*types.OrderItem, error) {
	stmt, err := s.db.Prepare("UPDATE order_items SET menu_id = ?, quantity = ?, price = ?, order_status = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	_, err = stmt.Exec(oiUpdate.MenuID, oiUpdate.Quantity, oiUpdate.Price, oiUpdate.OrderStatus, id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no order item found")
	}
	oi, err := s.GetOrderItemByID(id)
	if err != nil {
		return nil, fmt.Errorf("no order item found")
	}

	return oi, err
}

// change order item status
func (s *Store) ChangeOrderItemStatus(id int) (*types.HTTPGeneralRes, error) {
	oi, err := s.GetOrderItemByID(id)
	if err != nil {
		return nil, fmt.Errorf("no order item found")
	}

	var oiStatus types.OrderStatus

	stmt, err := s.db.Prepare("UPDATE order_items SET order_status = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	if oi.OrderStatus == "pending" {
		oiStatus = "completed"
	} else {
		oiStatus = "pending"
	}

	_, err = stmt.Exec(oiStatus, id)
	if err != nil {
		return nil, fmt.Errorf("no order item found")
	}

	orderitems, err := s.GetOrderItemByOrderID(oi.OrderID)
	if err != nil {
		return nil, fmt.Errorf("no order item found")
	}

	var isAllOrderItemsCompleted bool
	for _, oi := range *orderitems {
		if oi.OrderStatus == "completed" {
			isAllOrderItemsCompleted = true
		} else {
			isAllOrderItemsCompleted = false
		}
	}

	if isAllOrderItemsCompleted {
		err := s.UpdateOrderStatus(oi.OrderID, "completed")
		if err != nil {
			return nil, fmt.Errorf("error when updating order status")
		}

	}

	return &types.HTTPGeneralRes{Message: "order item status updated successfully", Success: true}, nil
}

func (s *Store) UpdateOrderStatus(id int, status string) error {
	stmt, err := s.db.Prepare("UPDATE orders SET order_status = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(status, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateOrder(oID, price, quantity int) error {
	order, err := s.GetOrderByID(oID)
	if err != nil {
		return err
	}
	stmt, err := s.db.Prepare("UPDATE orders SET total_price = ?, total_quantity = ? WHERE id = ?")
	if err != nil {
		return err
	}
	totalPrice := order.TotalPrice + (price * quantity)
	totalQuantity := order.TotalQuantity + quantity
	_, err = stmt.Exec(totalPrice, totalQuantity, oID)
	if err != nil {
		return err
	}
	return nil
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
