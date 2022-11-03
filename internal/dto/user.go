package dto

type UserCreateDto struct {
	FirstName string `json:"first_name" mapstructure:"first_name"`
	LastName  string `json:"last_name" mapstructure:"last_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
