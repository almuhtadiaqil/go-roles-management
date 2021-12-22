package web

type RoleUpdateRequest struct {
	Id          uint   `validate:"required" json:"id,string,omitempty"`
	Name        string `validate:"required,max=150,min=1" json:"name"`
	Label       string `validate:"required,max=20" json:"label"`
	Description string `validate:"required,max=250" json:"description"`
}
