package administratorRequest

import "time"

type UserRequest struct {
	IDRole              uint      `json:"id_role"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	RefreshToken        string    `json:"refresh_token"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired"`
}

type UserUpdateRequest struct {
	IDRole   uint   `json:"id_role"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
