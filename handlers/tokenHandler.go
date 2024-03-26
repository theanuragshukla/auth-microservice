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

type TokenResponse struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   data.Tokens `json:"data,omitempty"`
}

func (res *TokenResponse) toJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(res)
	return err
}
func (auth *Provider) TokenHandler(w http.ResponseWriter, r *http.Request) {
	RefreshToken := "x-refresh-token"
	Uid := "uid"
	uid := r.URL.Query().Get(Uid)
	refreshToken := r.Header.Get(RefreshToken)

	reqID := middlewares.GetTraceID(r)
	auth.L.Info("/token", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr), zap.String("uid", uid))

	response := TokenResponse{
		Status: false,
		Msg:    "Unable to generate access token",
	}

	if len(refreshToken) == 0 || len(uid) == 0 {
		auth.L.Info("null token or uid", zap.String("traceId", reqID), zap.String("uid", uid), zap.String("refresh", refreshToken))
		response.toJSON(w)
		return
	} else {
		session := data.Session{}
		claims := data.Claims{}
		token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			auth.L.Info(err.Error(), zap.String("traceId", reqID))
			response.Msg = "Unable to parse refresh token"
			response.toJSON(w)
			return
		}
		if claims.Subject != "refresh" {
			auth.L.Info("not refreshToken", zap.String("traceId", reqID))
			response.Msg = "Invalid refresh token"
			response.toJSON(w)
			return
		}
		auth.Db.DB.Where("uid = ?", claims.Uid).First(&session)
		if session.Seed == claims.Seed && session.Uid == claims.Uid {
			dbUser := data.User{}
			auth.Db.DB.Where("uid = ?", claims.Uid).First(&dbUser)
			if dbUser.Uid == session.Uid {
				tokens, err := GenerateTokens(*auth, dbUser)
				if err != nil {
					auth.L.Info("token error", zap.String("traceId", reqID))
					response.toJSON(w)
					return
				}
				auth.L.Info("token success", zap.String("traceId", reqID))
				response.Status = true
				response.Msg = "token generated"
				response.Data = tokens
				response.toJSON(w)
				return
			}
			auth.L.Info("uid not in Db", zap.String("traceId", reqID))
		} else {
			auth.L.Info("mismatch seed", zap.String("traceId", reqID))
		}
		response.toJSON(w)
	}
}
