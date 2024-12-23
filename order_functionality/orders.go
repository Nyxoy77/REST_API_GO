package orderfunctionality

type Order struct {
	OrderID   int    `json:"id" validate:"required"`
	UserID    int    `json:"user_id" validate:"required"`
	OrderDate string `json:"order_date" validate:"required"`
	Status    string `json:"status" validate:"required,oneof=pending shipped delivered canceled"`
}

type OrderItem struct {
	OrderItemID int     `json:"id" validate:"required"`
	OrderID     int     `json:"order_id" validate:"required"`
	ProductID   int     `json:"product_id" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required,min=1"`
	PriceAtTime float64 `json:"price_at_time" validate:"required,gt=0"`
}
