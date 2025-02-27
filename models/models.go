package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	FirstName string    `json:"firstname" validate:"required,min=2,max=100"`
	LastName  string    `json:"lastname" validate:"required,min=2,max=100"`
	Email     string    `json:"email" validate:"email,required"`
	Password  string    `json:"password" validate:"required,min=6"`
	Phone     string    `json:"phone" validate:"required,min=10,max=10"`
	CreatedAt time.Time `json:"created_at"`
	UserType  string    `json:"user_type" validate:"required,oneof=ADMIN USER"`
	UpdatedAt string    `json:"updated_at"`
	User_ID   int       `json:"id"`
	// Role      string    `json:"role" validate:"required,oneof=ADMIN USER"`
}

type Login struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type Forgot struct {
	Email string `json:"email"`
}

type Reset struct {
	User_ID     int8   `json:"user_id" validate:"required"`
	Expire_Time string `json:"expiration" validate:"required"`
	Reset_token string `json:"reset_token" validate:"required"`
	Used        bool   `json:"used" validate:"required"`
}

type UpdatePass struct {
	Password string `json:"password"`
}

type Product struct {
	Name          string  `json:"name" validate:"required,min=3,max=100"`
	Description   string  `json:"description" validate:"required,max=500"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	StockQuantity int     `json:"stock_quantity" validate:"required,min=0"`
	Status        string  `json:"status" validate:"required,oneof=active inactive discontinued"`
	Manufacturer  string  `json:"manufacturer" validate:"max=100"`
	ImageURL      string  `json:"image_url" validate:"url"`
}

type Claims struct {
	User_ID  int    `json:"user_id"`
	Email    string `json:"email"`
	UserType string `json:"user_type" validate:"required,oneof=ADMIN USER"`

	jwt.StandardClaims
}
