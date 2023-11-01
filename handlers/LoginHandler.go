package handlers

import (
	"auth-ms/data"
	"auth-ms/middlewares"
	"auth-ms/utils"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

func handleLoginError(err string, w io.Writer) {
	response := LoginResponse{
		Status: false,
		Msg:    err,
	}
	response.toJSON(w)
}

func (auth *Provider) LoginHandler(w http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetTraceID(r)
	auth.l.Info("/login", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr))
	var req LoginRequest
	var response LoginResponse
	err := req.fromJSON(r.Body)
	if err != nil {
		auth.l.Info("Invalid body", zap.String("traceId", reqID), zap.Int("status", http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		handleLoginError("Invalid data! Unable to parse the payload", w)
		return
	}

	err = auth.validate.Struct(req)
	errors := utils.NewValidationError()
	if err != nil {
		auth.l.Info("Validation error", zap.String("traceId", reqID), zap.Int("status", http.StatusBadRequest))
		for _, err := range err.(validator.ValidationErrors) {
			var el utils.ValidateErrorFormat
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			*errors = append(*errors, &el)
		}
		response = LoginResponse{
			Status: false,
			Msg:    "Validation Error",
			Errors: *errors,
		}
		response.toJSON(w)
		return
	}

	var dbUser data.User
	err = auth.db.DB.Where("email = ?", req.Email).First(&data.User{}).Scan(&dbUser).Error
	if err != nil {
		auth.l.Info("email not in db", zap.String("traceId", reqID), zap.Int("status", http.StatusUnauthorized))
		w.WriteHeader(http.StatusUnauthorized)
		handleLoginError("Wrong username or password", w)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		auth.l.Info("wrong pass", zap.String("traceId", reqID), zap.Int("status", http.StatusUnauthorized))
		handleLoginError("Wrong username or password", w)
		return
	}
	tokens, err := GenerateTokens(*auth, dbUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		auth.l.Info("token error", zap.String("traceId", reqID), zap.Int("status", http.StatusUnauthorized))
		handleLoginError("Unexpected server error", w)
		return
	}
	response = LoginResponse{
		Status: true,
		Msg:    "success",
		Data:   tokens,
	}
	auth.l.Info("login success", zap.String("traceId", reqID), zap.String("uid", dbUser.Uid), zap.Int("status", http.StatusOK))
	response.toJSON(w)
}
