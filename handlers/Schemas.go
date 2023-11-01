package handlers

import (
	"auth-ms/data"
	"auth-ms/utils"
	"encoding/json"
	"io"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

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
