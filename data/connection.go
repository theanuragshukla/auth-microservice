package data

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
)

func GetDb() (*gorm.DB, error) {
	dsn := url.URL{
		User:     url.UserPassword(viper.GetString("DB_USER"), viper.GetString("DB_PASS")),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", viper.Get("DB_HOST"), viper.GetInt("DB_PORT")),
		Path:     viper.GetString("DB_NAME"),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return db, nil
}
