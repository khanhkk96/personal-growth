package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"personal-growth/utils"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	// database
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	// server
	Env        string `mapstructure:"ENV"`
	ServerPort string `mapstructure:"PORT"`
	// authentication
	TokenSecret        string `mapstructure:"TOKEN_SECRET"`
	TokenExpiredIn     string `mapstructure:"TOKEN_EXPIRED_IN"`
	RefreshTokenMaxAge string `mapstructure:"REFRESH_TOKEN_MAX_AGE"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	// email
	EmailAddress  string `mapstructure:"EMAIL_ADDRESS"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`

	// redis
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`

	ClientAddress string `mapstructure:"CLIENT_ADDRESS"`
	ServerAddress string `mapstructure:"API_SERVER_ADDRESS"`
	ServerIP      string `mapstructure:"SERVER_IP"`

	// momo
	MomoMUrl        string `mapstructure:"MOMO_URL"`
	MomoPartnerCode string `mapstructure:"MOMO_PARTNER_CODE"`
	MomoAccessKey   string `mapstructure:"MOMO_ACCESS_KEY"`
	MomoSecretKey   string `mapstructure:"MOMO_SECRET_KEY"`

	// vnpay
	VnpTmnCode    string `mapstructure:"VNP_TMN_CODE"`
	VnpHashSecret string `mapstructure:"VNP_HASH_SECRET"`
	VnpUrl        string `mapstructure:"VNP_URL"`
	VnpVersion    string `mapstructure:"VNP_VERSION"`
}

func LoadConfig(path string) (config Config, err error) {
	env := os.Getenv("ENV")
	if utils.IsEmpty(&env) {
		env = "dev"
	}

	_, ferr := os.Stat(fmt.Sprintf("%s.env", env))
	if ferr != nil {
		typ := reflect.TypeOf(Config{})

		var configObj map[string]interface{} = make(map[string]interface{})
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i) // Lấy metadata field

			configObj[field.Name] = os.Getenv(field.Tag.Get("mapstructure"))
		}

		data, _ := json.Marshal(configObj)
		json.Unmarshal(data, &config)

		viper.AutomaticEnv()

		return
	}

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(env)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
