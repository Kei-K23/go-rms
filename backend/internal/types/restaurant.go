package types

type RestaurantStore interface {
	CreateRestaurant(r CreateRestaurant) (*Restaurant, error)
	GetRestaurantByID(rID int, accessToken string) (*Restaurant, error)
	DeleteRestaurant(rID int, accessToken string) (*HTTPGeneralRes, error)
	UpdateRestaurant(r UpdateRestaurant, rAccessToken string, rID int) (*Restaurant, error)
}

type Restaurant struct {
	ID          int    `json:id`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	OpenHours   string `json:"open_hours"`
	CloseHours  string `json:"close_hours"`
	CuisineType string `json:"cuisine_type"`
	AccessToken string `json:"access_token"`
	UserID      int    `json:"user_id"`
	Capacity    int    `json:"capacity"`
	CreatedAt   string `json:"created_at"`
}

type CreateRestaurant struct {
	Name        string `json:"name" validate="required"`
	Address     string `json:"address" validate="required"`
	Phone       string `json:"phone" validate="required"`
	OpenHours   string `json:"open_hours" validate="required"`
	CloseHours  string `json:"close_hours" validate="required"`
	CuisineType string `json:"cuisine_type" validate="required"`
	AccessToken string `json:"access_token" validate="required"`
	UserID      int    `json:"user_id" validate="required"`
	Capacity    int    `json:"capacity" validate="required"`
}

type UpdateRestaurant struct {
	Name        string `json:"name" validate="required"`
	Address     string `json:"address" validate="required"`
	Phone       string `json:"phone" validate="required"`
	OpenHours   string `json:"open_hours" validate="required"`
	CloseHours  string `json:"close_hours" validate="required"`
	CuisineType string `json:"cuisine_type" validate="required"`
	Capacity    int    `json:"capacity" validate="required"`
}
