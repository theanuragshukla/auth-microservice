package utils

import (
	"crypto/rand"
	"encoding/base64"
	"gorm.io/gorm"
)

func GenerateUID(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:len]
}

type Repository struct {
	DB *gorm.DB
}
