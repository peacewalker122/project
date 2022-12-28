package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string        `mapstructure:"DB_DRIVER"`
	DBSource          string        `mapstructure:"DB_SOURCE"`
	NotifDBSource     string        `mapstructure:"NOTIFDB_SOURCE"`
	RedisSource       string        `mapstructure:"REDIS_ADDR"`
	Email             string        `mapstructure:"EMAIL_HOST"`
	EmailPass         string        `mapstructure:"EMAIL_PASSWORD"`
	HTTPServerAddress string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AuthErrorAddres   string        `mapstructure:"REDIRECT_AUTH_TOKEN"`
	BucketAccount     string        `mapstructure:"GC_BUCKET_PROFILE"`
	BucketPost        string        `mapstructure:"GC_BUCKET_POST"`
	ClientOption      string        `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
	TokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshToken      time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
