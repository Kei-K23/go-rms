package types

type UserStore interface {
	CreateUser(user RegisterUser) (*RegisterUser, error)
	GetUserByEmail(user LoginUser) (*User, error)
	GetUserById(uID int) (*User, error)
	UpdateUser(user UpdateUser, uID int) (*User, error)
	DeleteUser(uID int) (*HTTPGeneralRes, error)
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	AccessKey string `json:"access_key"`
	CreatedAt string `json:"created_at"`
}

type UpdateUser struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}
