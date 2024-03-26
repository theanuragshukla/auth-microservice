package server

import (
	"auth-ms/auth-ms-grpc/protos/pb"
	"auth-ms/data"
	"auth-ms/handlers"
	"auth-ms/utils"
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthGrpcServer struct {
	pb.AuthServer
	auth *handlers.Provider
}

func NewAuthGrpcServer(auth *handlers.Provider) *AuthGrpcServer {
	return &AuthGrpcServer{auth: auth}
}

func (a *AuthGrpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	reqID := "grpc"
	a.auth.L.Info("/login", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr))

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

func (a *AuthGrpcServer) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	return &pb.SignupResponse{
		Status: true,
		Msg:    "",
		Data: &pb.Tokens{
			AccessToken:  "rthw45hwerh",
			RefreshToken: "werhtw45h",
			Uid:          "rthw45hw",
		},
		Errors: nil,
	}, nil
}
