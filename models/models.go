package models

import "time"

type User struct {
	FirstName     string    `json:"firstname" validate:"required,min=2,max=100"`
	LastName      string    `json:"lastname" validate:"required,min=2,max=100"`
	Email         string    `json:"email" validate:"email ,required"`
	Password      string    `json:"password" validate:"required,min = 6"`
	Phone         string    `json:"phone" validate:"required min=10,max=10"`
	Created_at    time.Time `json:"created_at"`
	User_type     string    `json:"user_type" validate:"required, eq=ADMIN | eq=USER"`
	Token         string    `json:"token"`
	Refresh_token string    `json:"refresh_token"`
	Updated_at    string    `json:"updated_at"`
	User_id       string    `json:"user_id"`
}
type Login struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}
