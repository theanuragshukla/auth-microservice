package handlers

import (
	"auth-ms/data"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
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
func (auth *AuthProvider) TokenHandler(w http.ResponseWriter, r *http.Request) {
	RefreshToken := "x-refresh-token"
	Uid := "uid"

	uid := r.URL.Query().Get(Uid)
	refreshToken := r.Header.Get(RefreshToken)

	response := TokenResponse{
		Status: false,
		Msg:    "Unable to generate access token",
	}

	if len(refreshToken) == 0 || len(uid) == 0 {
		response.toJSON(w)
		return
	} else {
		session := data.Session{}
		claims := data.Claims{}
		token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			response.Msg = "Unable to parse refresh token"
			response.toJSON(w)
			return
		}
		if claims.Subject != "refresh" {
			response.Msg = "Invalid refresh token"
			response.toJSON(w)
			return
		}
		auth.db.DB.Where("uid = ?", claims.Uid).First(&session)
		if session.Seed == claims.Seed && session.Uid == claims.Uid {
			dbUser := data.User{}
			auth.db.DB.Where("uid = ?", claims.Uid).First(&dbUser)
			if dbUser.Uid == session.Uid {
				tokens, err := GenerateTokens(*auth, dbUser)
				if err != nil {
					response.toJSON(w)
					return
				}
				response.Status = true
				response.Msg = "token generated"
				response.Data = tokens
				response.toJSON(w)
				return
			}
		}
		response.toJSON(w)
	}
}
