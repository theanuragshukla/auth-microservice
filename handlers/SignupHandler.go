package handlers

import (
	"auth-ms/data"
	"auth-ms/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

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

func handleSignupError(err string, w io.Writer) {
	response := SignupResponse{
		Status: false,
		Msg:    err,
	}
	response.toJSON(w)
}
func (auth *AuthProvider) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user data.User
	var response SignupResponse
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handleSignupError("Invalid payload format! Unable to unmarshal", w)
		return
	}
	err = auth.validate.Struct(user)
	errors := utils.NewValidationError()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el utils.ValidateErrorFormat
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			*errors = append(*errors, &el)
		}
		response = SignupResponse{
			Status: false,
			Msg:    "Validation Error",
			Errors: *errors,
		}
		response.toJSON(w)
		return
	}
	user.Uid = utils.GenerateUID(32)
	pass := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handleSignupError("Unable to process the request", w)
		return
	}
	user.Password = string(hashedPassword)
	res := auth.db.DB.Create(&user)
	if res.Error != nil {
		w.WriteHeader(http.StatusConflict)
		handleSignupError("Email address already in use", w)
		return
	}

	tokens, err := GenerateTokens(*auth, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handleLoginError("Unexpected server error", w)
		return
	}

	ret := SignupResponse{
		Status: true,
		Msg:    "success",
		Data:   tokens,
	}
	ret.toJSON(w)
}
