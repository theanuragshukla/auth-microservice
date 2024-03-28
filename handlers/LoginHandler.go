package handlers

import (
	"io"
	"net/http"

	"auth-ms/data"
	"auth-ms/middlewares"
	"auth-ms/utils"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func handleLoginError(err string, w io.Writer) {
	response := LoginResponse{
		Status: false,
		Msg:    err,
	}
	response.toJSON(w)
}

// swagger:parameters login
type LoginRequestParams struct {
	// Request body containing email and password
	//
	// in: body
	Body LoginRequest
}

// swagger:route POST /login login
// Returns Tokens and uid if the credentials provided are correct
// responses:
// 200: LoginResponse

func (auth *Provider) LoginHandler(w http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetTraceID(r)
	auth.L.Info("/login", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr))
	var req LoginRequest
	var response LoginResponse
	err := req.fromJSON(r.Body)
	if err != nil {
		auth.L.Info("Invalid body", zap.String("traceId", reqID), zap.Int("status", http.StatusBadRequest))
		w.WriteHeader(http.StatusBadRequest)
		handleLoginError("Invalid data! Unable to parse the payload", w)
		return
	}

	err = auth.Validate.Struct(req)
	errors := utils.NewValidationError()
	if err != nil {
		auth.L.Info("Validation error", zap.String("traceId", reqID), zap.Int("status", http.StatusBadRequest))
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
	err = auth.Db.DB.Where("email = ?", req.Email).First(&data.User{}).Scan(&dbUser).Error
	if err != nil {
		auth.L.Info("email not in Db", zap.String("traceId", reqID), zap.Int("status", http.StatusUnauthorized))
		w.WriteHeader(http.StatusUnauthorized)
		handleLoginError("Wrong username or password", w)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		auth.L.Info("wrong pass", zap.String("traceId", reqID), zap.Int("status", http.StatusUnauthorized))
		handleLoginError("Wrong username or password", w)
		return
	}
	tokens, err := GenerateTokens(*auth, dbUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		auth.L.Info("token error", zap.String("traceId", reqID), zap.Int("status", http.StatusUnauthorized))
		handleLoginError("Unexpected server error", w)
		return
	}
	response = LoginResponse{
		Status: true,
		Msg:    "success",
		Data:   tokens,
	}
	auth.L.Info("login success", zap.String("traceId", reqID), zap.String("uid", dbUser.Uid), zap.Int("status", http.StatusOK))
	response.toJSON(w)
}
