package types

type RestaurantTablesStore interface {
	CreateRestaurantTable(rT CreateRestaurantTable) (*RestaurantTable, error)
}

type Status string

const (
	Occupied Status = "Occupied"
	Vacant   Status = "Vacant"
)

type RestaurantTable struct {
	ID           int    `json:id`
	TableNumber  int    `json:"table_number"`
	Status       Status `json:"status"`
	Capacity     int    `json:"capacity"`
	RestaurantID int    `json:"restaurant_id"`
	CreatedAt    string `json:"created_at"`
}

type CreateRestaurantTable struct {
	TableNumber  int    `json:"table_number" validate:"required"`
	Status       Status `json:"status" validate:"required,oneof=occupied vacant"`
	Capacity     int    `json:"capacity" validate:"required"`
	RestaurantID int    `json:"restaurant_id" validate:"required"`
}

type UpdateRestaurantTable struct {
	TableNumber int    `json:"table_number" validate:"required"`
	Status      Status `json:"status" validate:"required,oneof=occupied vacant"`
	Capacity    int    `json:"capacity" validate:"required"`
}