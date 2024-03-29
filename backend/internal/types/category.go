package types

type CategoryStore interface {
	CreateCategory(ct CreateCategory) (*Category, error)
	GetCategories() (*[]Category, error)
	GetCategoryByID(id int) (*Category, error)
}

type Category struct {
	ID          int    `json:id`
	Name        string `json:name`
	Description string `json:description`
	CreatedAt   string `json:"created_at"`
}

type CreateCategory struct {
	Name        string `json:name`
	Description string `json:description`
}

type UpdateCategory struct {
	Name        string `json:name`
	Description string `json:description`
}