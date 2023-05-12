package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// SignUpReq represent signup request body
type SignUpReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request SignUpReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Username, validation.Required, validation.Length(6, 30)),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required, validation.Length(8, 20)),
	)
}

// SignInReq represent signin request body
type LogInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request LogInReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required),
	)
}
