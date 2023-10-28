package data

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	Uid       string    `gorm:"not null;size:32" json:"-"`
	FirstName string    `gorm:"not null;size:50" json:"firstName"`
	LastName  string    `gorm:"size:50" json:"lastName,omitempty"`
	Email     string    `gorm:"not null; unique" json:"email"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `sql:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `sql:"autoUpdateTime" json:"-"`
}

type Session struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	Uid       string `gorm:"not null;size:32;unique" json:"-"`
	Seed      string
	StartedAt time.Time `sql:"autoCreateTime" json:"-"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	UserId       string `json:"userId,omitempty"`
}

func (t *Tokens) toJSON() ([]byte, error) {
	enc, err := json.Marshal(t)
	return enc, err
}

type Claims struct {
	Uid  string
	Seed string
	jwt.StandardClaims
}

func GetClaims(uid string, seed string, subject string, expiresIn time.Duration) Claims {
	claims := Claims{
		Uid:  uid,
		Seed: seed,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresIn).UnixMilli(),
			Subject:   subject,
		},
	}

	return claims
}
