package types

type UserStore interface {
	CreateUser(user RegisterUser) (*RegisterUser, error)
	GetUserByEmail(user LoginUser) (*User, error)
	GetUserById(uID int) (*User, error)
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
