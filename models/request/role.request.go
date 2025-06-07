package request


type Roles struct {
	NameRole string `json:"name_role" validate:"required,min=3"`
}