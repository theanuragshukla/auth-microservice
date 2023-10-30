package utils

type Config struct {
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort int8   `mapstructure:"DB_PORT"`
	DBName string `mapstructure:"DB_NAME"`
}
