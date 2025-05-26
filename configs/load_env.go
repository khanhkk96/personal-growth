package configs

import (
	"encoding/json"
	"os"
	"reflect"
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
	RefreshTokenMaxAge string        `mapstructure:"REFRESH_TOKEN_MAX_AGE"`
	RefreshTokenSecret string        `mapstructure:"REFRESH_TOKEN_SECRET"`
	//email
	EmailAddress  string `mapstructure:"EMAIL_ADDRESS"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
	ClientAddress string `mapstructure:"CLIENT_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	_, ferr := os.Stat("dev.env")
	if ferr != nil {
		typ := reflect.TypeOf(Config{})

		var configObj map[string]interface{} = make(map[string]interface{})
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i) // Lấy metadata field

			configObj[field.Name] = os.Getenv(field.Tag.Get("mapstructure"))
		}

		data, _ := json.Marshal(configObj)
		json.Unmarshal(data, &config)

		return
	}

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
