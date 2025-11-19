package models

import "time"

type User struct {
	IDUser              uint      `json:"id_user"`
	IDRole              uint      `json:"id_role"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	RefreshToken        string    `json:"refresh_token"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired"`
	IsActive            bool      `json:"is_active"`
	CreatedDate         time.Time `json:"created_date"`
	UpdatedDate         time.Time `json:"updated_date"`
}
