package models

import (
	"time"
)

type User struct {
	FirstName    string    `json:"firstname" validate:"required,min=2,max=100"`
	LastName     string    `json:"lastname" validate:"required,min=2,max=100"`
	Email        string    `json:"email" validate:"email,required"`
	Password     string    `json:"password" validate:"required,min=6"`
	Phone        string    `json:"phone" validate:"required,min=10,max=10"`
	CreatedAt    time.Time `json:"created_at"`
	UserType     string    `json:"user_type" validate:"required,oneof=ADMIN USER"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	UpdatedAt    string    `json:"updated_at"`
	User_ID      int8      `json:"id"`
}

type Login struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type Forgot struct {
	Email string `json:"email"`
}

type Reset struct {
	User_ID     int8   `json:"user_id"`
	Expire_Time string `json:"expiration"`
	Reset_token string `json:"reset_token"`
	Used        bool   `json:"used"`
}

type UpdatePass struct {
	Password string `json:"password" validate:"required,min=6"`
}
