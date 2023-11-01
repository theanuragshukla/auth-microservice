package handlers

import (
	"auth-ms/data"
	"auth-ms/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"time"
)

type Provider struct {
	db       *utils.Repository
	validate *validator.Validate
	l        *zap.Logger
}

func NewProvider(repo *utils.Repository, logger *zap.Logger) *Provider {
	return &Provider{repo, validator.New(), logger}
}

func (auth *Provider) SaveSession(uid, seed string) error {
	res := auth.db.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}},
		DoUpdates: clause.AssignmentColumns([]string{"seed", "started_at"}),
	}).Create(&data.Session{
		Uid:  uid,
		Seed: seed,
	})
	return res.Error
}
func GenerateTokens(auth Provider, user data.User) (data.Tokens, error) {
	seed := utils.GenerateUID(16)
	var tokens data.Tokens
	accessExpireTime := 5 * time.Minute
	refreshExpireTime := 30 * 24 * 60 * time.Minute

	accessClaims := data.GetClaims(user.Uid, seed, "access", accessExpireTime)
	refreshClaims := data.GetClaims(user.Uid, seed, "refresh", refreshExpireTime)

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessStr, _ := access.SignedString([]byte(viper.GetString("JWT_SECRET")))
	refreshStr, _ := refresh.SignedString([]byte(viper.GetString("JWT_SECRET")))
	tokens = data.Tokens{
		AccessToken: accessStr, RefreshToken: refreshStr, UserId: user.Uid,
	}
	err := auth.SaveSession(user.Uid, seed)
	return tokens, err
}
