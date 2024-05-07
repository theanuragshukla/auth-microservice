package handlers

import (
	"encoding/json"
	"io"

	"auth-ms/data"
	"auth-ms/utils"
)

// swagger:parameters LoginRequest
// LoginRequest is the request model for the login endpoint
type LoginRequest struct {
	Email    string `json:"email" Validate:"required,email"`
	Password string `json:"password" Validate:"required"`
}

// swagger:response LoginResponse
// LoginResponse is the response model for the login endpoint
type LoginResponse struct {
	Status bool                `json:"status"`
	Msg    string              `json:"msg"`
	Data   data.Tokens         `json:"data,omitempty"`
	Errors utils.ValidateError `json:"errors,omitempty"`
}

func (req *LoginResponse) toJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(req)
}

func (req *LoginRequest) fromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(&req)
	return err
}

// swagger:model
// SignupRequest is the request model for the signup endpoint
type SignupRequest struct {
	FirstName string `gorm:"not null;size:50" json:"firstName" validate:"required,min=3,max=50"`
	LastName  string `gorm:"size:50" json:"lastName,omitempty" validate:"max=50"`
	Email     string `gorm:"not null; unique" json:"email" validate:"required,email"`
	Password  string `gorm:"not null" validate:"required,min=8,max=128" json:"password"`
}

func (req *SignupRequest) fromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(&req)
	return err
	}

// swagger:response SignupResponse
// SignupResponse is the response model for the signup endpoint
type SignupResponse struct {
	Status bool                `json:"status" `
	Msg    string              `json:"msg,omitempty"`
	Data   data.Tokens         `json:"data,omitempty"`
	Errors utils.ValidateError `json:"errors,omitempty"`
}

func (res *SignupResponse) toJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(&res)
}
