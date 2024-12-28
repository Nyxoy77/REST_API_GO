package user_customer

import "time"

type ItemToBeAdded struct {
	User_Id    int `json:"user_id" validate:"required"`
	Product_Id int `json:"product_id" validate:"required"`
	Quantity   int `json:"quantity" validate:"required"`
}

type CartItem struct {
	Id             int     `json:"id,omitempty"`
	User_Id        int     `json:"user_id" validate:"required"`
	Product_Id     int     `json:"product_id" validate:"required"`
	Quantity       int     `json:"quantity" validate:"required,min=1,max=7"`
	Price_Per_Unit float64 `json:"price_per_unit" validate:"required"`
}

type DeleteItem struct {
	User_Id    int `json:"user_id" validate:"required"`
	Product_Id int `json:"product_id" validate:"required"`
	Quantity   int `json:"quantity" validate:"required,min=1,max=7"`
}

type Orders struct {
	Id         int       `json:"id,omitempty"`
	User_Id    int       `json:"user_id" validate:"required"`
	Order_Date time.Time `json:"order_date" validate:"required"`
	Status     string    `json:"status" validate:"required,oneof=pending shipped delivered shipped"`
}

type OrderItems struct {
	Id            int     `json:"id,omitempty"`
	Order_Id      int     `json:"order_id" validate:"required"`
	Product_Id    int     `json:"product_id" validate:"required"`
	Quantity      int     `json:"quantity" validate:"required,min=1,max=7"`
	Price_At_Time float64 `json:"price_at_time" validate:"required"`
}

type FrontEndOrder struct {
	Product_Id int `json:"product_id" validate:"required"`
	Quantity   int `json:"quantity" validate:"required,min=1,max=7"`
}
type CombinedOrder struct {
	User_Id int             `json:"user_id" validate:"required"`
	Items   []FrontEndOrder `json:"items" validate:"required"`
}
