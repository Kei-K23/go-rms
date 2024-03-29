package types

type MenuStore interface {
	CreateMenu(m CreateMenu, rID int) (*Menu, error)
	GetMenuByID(id int) (*Menu, error)
	GetMenuByRestaurantID(rID int) (*[]Menu, error)
}

type Menu struct {
	ID           int    `json:id`
	Name         string `json:name`
	Description  string `json:description`
	Available    bool   `json:available`
	CategoryID   int    `json:category_id`
	RestaurantID int    `json:restaurant_id`
	Price        int    `json:price`
	CreatedAt    string `json:created_at`
}

type CreateMenu struct {
	Name        string `json:name`
	Description string `json:description`
	Available   bool   `json:available`
	CategoryID  int    `json:"category_id"`
	Price       int    `json:price`
}

type UpdateMenu struct {
	Name        string `json:name`
	Description string `json:description`
	Available   bool   `json:available`
	CategoryID  int    `json:category_id`
	Price       int    `json:price`
}
