package administratorResponse

import "time"

type RoleResponse struct {
	IDRole      uint      `json:"id_role"`
	RoleName    string    `json:"role_name"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
