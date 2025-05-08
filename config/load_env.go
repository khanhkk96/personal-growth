package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	//database
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	//server
	ServerPort string `mapstructure:"PORT"`
	//authentication
	TokenSecret        string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiredIn     time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge        int           `mapstructure:"TOKEN_MAX_AGE"`
	RefreshTokenSecret string        `mapstructure:"REFRESH_TOKEN_SECRET"`
	//email
	EmailAddress  string `mapstructure:"EMAIL_ADDRESS"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("dev")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
