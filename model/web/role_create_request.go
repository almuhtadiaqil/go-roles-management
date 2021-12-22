package web

type RoleCreateRequest struct {
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Label       string `validate:"required,max=20" json:"label"`
	Description string `validate:"required,max=250" json:"description"`
}
