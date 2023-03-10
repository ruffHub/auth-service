package usermodel

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Job      string `json:"job,omitempty" validate:"required"`
}
