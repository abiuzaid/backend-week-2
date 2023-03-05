package models

import "time"

type Cart struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int                 `json:"user_id"`
	User          UsersCartResponse   `json:"user"`
	ProductID     int                 `json:"product_id" gorm:"type: int"`
	Product       ProductCartResponse `json:"product"`
	OrderQuantity int                 `json:"order_quantity" gorm:"type: int"`
	CreatedAt     time.Time           `json:"-"`
	UpdatedAt     time.Time           `json:"-"`
}

type CartUSerResponse struct {
	ProductID     int `json:"product_id" gorm:"type: int"`
	OrderQuantity int `json:"order_quantity" gorm:"type: int"`
	UserID        int `json:"-"`
}

type CartProductResponse struct {
	ProductID     int `json:"-"`
	OrderQuantity int `json:"order_quantity" gorm:"type: int"`
	UserID        int `json:"user_id" gorm:"type: int"`
}

func (CartUSerResponse) TableName() string {
	return "carts"
}

func (CartProductResponse) TableName() string {
	return "carts"
}
