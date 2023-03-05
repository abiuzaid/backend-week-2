package models

import "time"

type Transaction struct {
	ID            int                        `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int                        `json:"user_id"`
	User          UsersTransactionResponse   `json:"user"`
	ProductID     int                        `json:"product_id" gorm:"type: int"`
	Product       ProductTransactionResponse `json:"product"`
	OrderQuantity int                        `json:"order_quantity" gorm:"type: int"`
	CreatedAt     time.Time                  `json:"-"`
	UpdatedAt     time.Time                  `json:"-"`
}

type TransactionUSerResponse struct {
	ProductID     int                        `json:"product_id" gorm:"type: int"`
	Product       ProductTransactionResponse `json:"product"`
	OrderQuantity int                        `json:"order_quantity" gorm:"type: int"`
	UserID        int                        `json:"-"`
}

type TransactionProductResponse struct {
	ProductID     int                      `json:"-"`
	UserID        int                      `json:"user_id" gorm:"type: int"`
	User          UsersTransactionResponse `json:"user"`
	OrderQuantity int                      `json:"order_quantity" gorm:"type: int"`
}

func (TransactionUSerResponse) TableName() string {
	return "transactions"
}

func (TransactionProductResponse) TableName() string {
	return "transactions"
}
