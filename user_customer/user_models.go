package user_customer

type Item struct {
	Id             int     `json:"id"`
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
