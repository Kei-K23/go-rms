package types

type AuthStore interface {
	HashedPassword(password string) (string, error)
	VerifyPassword(hPassword, password string) error
}

type RegisterUser struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=30"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	AccessKey string `json:"access_key" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}
