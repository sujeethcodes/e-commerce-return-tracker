package entity

type ReturnProducts struct {
	ProductID  string `json:"product_id" validate:"required"`
	OrderID    string `json:"order_id" validate:"required"`
	AccountID  string `json:"account_id" validate:"required"`
	Reason     string `json:"reason" validate:"required"`
	ReturnDate string `json:"return_date" validate:"required"`
}
