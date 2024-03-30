package types

type OrderItemStore interface {
	CreateOrderItem(orderItem CreateOrderItem) (*OrderItem, error)
	ChangeOrderItemStatus(id int) (*HTTPGeneralRes, error)
	UpdateOrderStatus(id int, status string) error
	UpdateOrder(oID, price, quantity int) error
}

type OrderItem struct {
	ID          int         `json:"id"`
	Price       int         `json:"price"`
	Quantity    int         `json:"quantity"`
	OrderID     int         `json:"order_id"`
	MenuID      int         `json:"menu_id"`
	OrderStatus OrderStatus `json:"order_status"`
	CreatedAt   string      `json:"created_at"`
}

type CreateOrderItem struct {
	Price       int         `json:"price" validate:"required"`
	Quantity    int         `json:"quantity" validate:"required"`
	OrderID     int         `json:"order_id"`
	MenuID      int         `json:"menu_id" validate:"required"`
	OrderStatus OrderStatus `json:"order_status" validate:"required"`
}

type UpdateOrderItem struct {
	Price       int         `json:"price" validate:"required"`
	Quantity    int         `json:"quantity" validate:"required"`
	MenuID      int         `json:"menu_id" validate:"required"`
	OrderStatus OrderStatus `json:"order_status" validate:"required"`
}
