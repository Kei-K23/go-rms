package types

type PaymentStore interface {
	CreatePayment(payment CreatePayment) error
}

type Payment struct {
	ID          int    `json:"id"`
	OrderID     int    `json:"order_id"`
	Amount      int    `json:"amount"`
	PaymentType string `json:"payment_type"`
	CreatedAt   string `json:"created_at"`
}

type CreatePayment struct {
	OrderID     int    `json:"order_id"`
	Amount      int    `json:"amount"`
	PaymentType string `json:"payment_type"`
}
