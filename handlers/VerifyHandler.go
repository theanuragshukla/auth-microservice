package handlers

import (
	"auth-ms/data"
	"auth-ms/middlewares"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type VerifyResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg,omitempty"`
}

func (res *VerifyResponse) toJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(res)
	return err
}

func (auth *Provider) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetTraceID(r)
	auth.L.Info("/verify", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr))
	AccessToken := "x-access-token"
	accessToken := r.Header.Get(AccessToken)
	response := VerifyResponse{
		Status: false,
		Msg:    "Access token not provided",
	}
	if len(accessToken) == 0 {
		auth.L.Info("token nil", zap.String("traceId", reqID))
		response.toJSON(w)
		return
	} else {
		session := data.Session{}
		claims := data.Claims{}
		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			auth.L.Info("invalid token", zap.String("traceId", reqID))
			response.Msg = "Unable to parse access token"
			response.toJSON(w)
			return
		}
		if claims.Subject != "access" {
			auth.L.Info("not accessToken", zap.String("traceId", reqID))
			response.Msg = "Invalid access token"
			response.toJSON(w)
			return
		}
		auth.Db.DB.Where("uid = ?", claims.Uid).First(&session)
		if session.Seed == claims.Seed {
			auth.L.Info("verified", zap.String("traceId", reqID), zap.String("uid", claims.Uid))
			response.Status = true
			response.Msg = "verified"
			response.toJSON(w)
			return
		} else {
			auth.L.Info("seed mismatch or nil", zap.String("traceId", reqID), zap.String("uid", claims.Uid))
		}
		response.Msg = "unverified"
		response.toJSON(w)
	}
}
