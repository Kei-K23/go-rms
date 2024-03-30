package types

type OrderStore interface {
	CreateOrder(order CreateOrder, orderItems []CreateOrderItem) (*Order, error)
	DeleteOrder(oID, rID int) (*HTTPGeneralRes, error)
}

type OrderStatus string

const (
	PENDING   OrderStatus = "pending"
	COMPLETED OrderStatus = "completed"
)

type Order struct {
	ID            int         `json:"id"`
	TableNumber   int         `json:"table_number"`
	TotalPrice    int         `json:"total_price"`
	TotalQuantity int         `json:"total_quantity"`
	OrderTime     string      `json:"order_time"`
	OrderStatus   OrderStatus `json:"order_status"`
	RestaurantId  int         `json:"restaurant_id"`
	CreatedAt     string      `json:"created_at"`
}

type CreateOrder struct {
	TableNumber   int         `json:"table_number" validate:"required"`
	TotalPrice    int         `json:"total_price" validate:"required"`
	TotalQuantity int         `json:"total_quantity" validate:required`
	OrderTime     string      `json:"order_time"`
	OrderStatus   OrderStatus `json:"order_status" validate:"required,oneof=pending completed"`
	RestaurantId  int         `json:"restaurant_id" validate:"required"`
}

type UpdateOrder struct {
	TableNumber int         `json:"table_number" validate:"required"`
	TotalPrice  int         `json:"total_price" validate:"required"`
	OrderTime   string      `json:"order_time"`
	OrderStatus OrderStatus `json:"order_status" validate:"required,oneof=pending completed"`
}
