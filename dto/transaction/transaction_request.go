package transactionsdto

// type TransactionRequest struct {
// 	Name    string `json:"name"`
// 	Email   string `json:"email"`
// 	Phone   string `json:"phone"`
// 	Address string `json:"address"`
// 	Total   int32  `json:"total"`
// }

// mas fauzan yg bawah

type TransactionRequest struct {
	ProductID     int `json:"product_id" gorm:"type: int"`
	OrderQuantity int `json:"order_quantity" gorm:"type: int" validate:"required"`
	UserID        int `json:"user_id"`
}
