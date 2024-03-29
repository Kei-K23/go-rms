package types

type RestaurantTablesStore interface {
	GetRestaurantTables(rID int) (*[]RestaurantTable, error)
	CreateRestaurantTable(rT CreateRestaurantTable) (*RestaurantTable, error)
	GetRestaurantTableByID(rTID, rID int) (*RestaurantTable, error)
	UpdateRestaurantTable(rT UpdateRestaurantTable, rID, rTID int) (*RestaurantTable, error)
	DeleteRestaurantTable(rID, rTID int) (*HTTPGeneralRes, error)
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
	Status   Status `json:"status" validate:"required,oneof=occupied vacant damage"`
	Capacity int    `json:"capacity" validate:"required"`
}
