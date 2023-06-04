package form

type UpdateUser struct {
	Name string `json:"name" form:"name"`
	Role string `json:"role" form:"role"`
}
