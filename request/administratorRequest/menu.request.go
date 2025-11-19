package administratorRequest

type MenuRequest struct {
	IdMenu     string `json:"id_menu"`
	IdParent   string `json:"id_parent"`
	NamaMenu   string `json:"nama_menu"`
	Controller string `json:"controller"`
	ClassIcon  string `json:"class_icon"`
	MenuSort   uint   `json:"menu_sort"`
	IsActive   bool   `json:"is_active"`
}
