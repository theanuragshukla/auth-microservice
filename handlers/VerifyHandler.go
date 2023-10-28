package handlers

import (
	"auth-ms/data"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
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

func (auth *AuthProvider) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	AccessToken := "x-access-token"
	accessToken := r.Header.Get(AccessToken)
	response := VerifyResponse{
		Status: false,
		Msg:    "Access token not provided",
	}
	if len(accessToken) == 0 {
		response.toJSON(w)
		return
	} else {
		session := data.Session{}
		claims := data.Claims{}
		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			response.Msg = "Unable to parse access token"
			response.toJSON(w)
			return
		}
		if claims.Subject != "access" {
			response.Msg = "Invalid access token"
			response.toJSON(w)
			return
		}
		auth.db.DB.Where("uid = ?", claims.Uid).First(&session)
		if session.Seed == claims.Seed {
			response.Status = true
			response.Msg = "verified"
			response.toJSON(w)
			return
		}
		response.Msg = "unverified"
		response.toJSON(w)
	}
}
