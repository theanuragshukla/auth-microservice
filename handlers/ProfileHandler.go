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

type ProfileResponse struct {
	Status bool      `json:"status"`
	Msg    string    `json:"msg,omitempty"`
	Data   data.User `json:"data,omitempty"`
}

func (res *ProfileResponse) toJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(res)
	return err
}

func (auth *Provider) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetTraceID(r)
	auth.l.Info("/profile", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr))
	AccessToken := "x-access-token"
	accessToken := r.Header.Get(AccessToken)
	response := ProfileResponse{
		Status: false,
		Msg:    "Access token not provided",
	}
	if len(accessToken) == 0 {
		auth.l.Info("token nil", zap.String("traceId", reqID))
		response.toJSON(w)
		return
	} else {
		session := data.Session{}
		claims := data.Claims{}
		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			auth.l.Info("invalid token", zap.String("traceId", reqID))
			response.Msg = "Unable to parse access token"
			response.toJSON(w)
			return
		}
		if claims.Subject != "access" {
			auth.l.Info("not accessToken", zap.String("traceId", reqID))
			response.Msg = "Invalid access token"
			response.toJSON(w)
			return
		}
		auth.db.DB.Where("uid = ?", claims.Uid).First(&session)
		if session.Seed == claims.Seed {
			var dbUser data.User
			err = auth.db.DB.Where("uid = ?", claims.Uid).First(&data.User{}).Scan(&dbUser).Error
			if err != nil {
				auth.l.Info("uid not in db", zap.String("traceId", reqID), zap.Int("status", http.StatusUnauthorized), zap.String("uid", claims.Uid))
				w.WriteHeader(http.StatusUnauthorized)
				handleLoginError("Wrong username or password", w)
				return
			}
			response.Status = true
			response.Msg = "verified"
			response.Data = dbUser
			response.toJSON(w)
			auth.l.Info("Profile returned", zap.String("traceId", reqID), zap.String("uid", claims.Uid))
			return
		} else {
			auth.l.Info("seed mismatch or nil", zap.String("traceId", reqID), zap.String("uid", claims.Uid))
		}
		response.Msg = "unverified"
		response.toJSON(w)
	}
}
