package authenticationRequest

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
	// Username string `json:"username" validate:"required"`
	Password string `json:"password" gorm:"unique" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" gorm:"unique" validate:"required"`
}
