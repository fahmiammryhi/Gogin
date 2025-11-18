package response

type Response struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
	Data       interface{} `json:"data"`
}

type ResponsePage struct {
	Page       uint `json:"page"`
	PageSize   uint `json:"page_size"`
	TotalItems uint `json:"total_items"`
	TotalPages uint `json:"total_pages"`
}
