package handlers

import (
	"auth-ms/data"
	"auth-ms/middlewares"
	"auth-ms/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

func handleSignupError(err string, w io.Writer) {
	response := SignupResponse{
		Status: false,
		Msg:    err,
	}
	response.toJSON(w)
}
func (auth *Provider) SignupHandler(w http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetTraceID(r)
	auth.l.Info("/signup", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr))
	var user data.User
	var response SignupResponse
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		auth.l.Info("Invalid body", zap.String("traceId", reqID), zap.Int("status", http.StatusBadRequest))
		handleSignupError("Invalid payload format! Unable to unmarshal", w)
		return
	}
	err = auth.validate.Struct(user)
	errors := utils.NewValidationError()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			auth.l.Info("Validation error", zap.String("traceId", reqID), zap.Int("status", http.StatusBadRequest))
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
		auth.l.Info("hash error", zap.String("traceId", reqID), zap.Int("status", http.StatusInternalServerError))
		w.WriteHeader(http.StatusInternalServerError)
		handleSignupError("Unable to process the request", w)
		return
	}
	user.Password = string(hashedPassword)
	res := auth.db.DB.Create(&user)
	if res.Error != nil {
		w.WriteHeader(http.StatusConflict)
		auth.l.Info("duplicate email", zap.String("traceId", reqID), zap.Int("status", http.StatusConflict))
		handleSignupError("Email address already in use", w)
		return
	}

	tokens, err := GenerateTokens(*auth, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		auth.l.Info("token error", zap.String("traceId", reqID), zap.Int("status", http.StatusInternalServerError))
		handleLoginError("Unexpected server error", w)
		return
	}

	ret := SignupResponse{
		Status: true,
		Msg:    "success",
		Data:   tokens,
	}
	auth.l.Info("signup success", zap.String("traceId", reqID), zap.Int("status", http.StatusOK))
	ret.toJSON(w)
}
