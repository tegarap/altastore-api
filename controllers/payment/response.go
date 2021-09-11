package payment

type Response struct {
	ID          uint   `gorm:"primarykey"`
	PaymentName string `json:"payment_name"`
}