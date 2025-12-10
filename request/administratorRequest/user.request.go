package administratorRequest

import "time"

type UserRequest struct {
	IDRole              uint      `json:"id_role" binding:"required"`
	Username            string    `json:"username" binding:"required,min=3"`
	Email               string    `json:"email" binding:"required,email"`
	Password            string    `json:"password" binding:"required,min=6"`
	RefreshToken        string    `json:"refresh_token"` // opsional
	RefreshTokenExpired time.Time `json:"refresh_token_expired"`
}

type UserUpdateRequest struct {
	IDRole   uint   `json:"id_role"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
