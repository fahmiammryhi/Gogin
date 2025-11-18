package models

type User struct {
	IDUser      int    `json:"id_user"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}
