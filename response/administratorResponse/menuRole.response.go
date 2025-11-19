package administratorResponse

type MenuRoleResponse struct {
	IdMenuRole uint   `json:"id_menu_role"`
	IdRole     uint   `json:"id_role"`
	IdMenu     string `json:"id_menu"`
	IsView     bool   `json:"is_view"`
	IsInsert   bool   `json:"is_insert"`
	IsEdit     bool   `json:"is_edit"`
	IsDelete   bool   `json:"is_delete"`
	IsPrint    bool   `json:"is_print"`
	IsApprove  bool   `json:"is_approve"`
	IsActive   bool   `json:"is_active"`
}
