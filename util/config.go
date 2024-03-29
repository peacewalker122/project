package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver           string        `mapstructure:"DB_DRIVER"`
	DBSource           string        `mapstructure:"DB_SOURCE"`
	RedisSource        string        `mapstructure:"REDIS_ADDR"`
	Email              string        `mapstructure:"EMAIL_HOST"`
	EmailPass          string        `mapstructure:"EMAIL_PASSWORD"`
	BaseURL            string        `mapstructure:"BASE_URL"`
	BucketName         string        `mapstructure:"GCP_BUCKET_NAME"`
	SignUpLink         string        `mapstructure:"SIGNUPLINK"`
	HTTPServerAddress  string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenKey           string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AuthErrorAddress   string        `mapstructure:"REDIRECT_AUTH_TOKEN"`
	BucketAccount      string        `mapstructure:"GC_BUCKET_PROFILE"`
	BucketPost         string        `mapstructure:"GC_BUCKET_POST"`
	ClientOption       string        `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
	GoogleClientID     string        `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleClientSecret string        `mapstructure:"GOOGLE_OAUTH_SECRET"`
	UserDir            string        `mapstructure:"USER_DIR"`
	EmailSMTP          string        `mapstructure:"EMAIL_SMTP"`
	MaxUploadSize      int64         `mapstructure:"MAX_UPLOAD_SIZE"`
	TokenDuration      time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshToken       time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("MAX_UPLOAD_SIZE", 30)

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func (c *Config) MaxUploadSizeInBytes() int64 {
	return c.MaxUploadSize * 1024 * 1024
}
