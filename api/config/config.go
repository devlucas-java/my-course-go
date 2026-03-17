package config

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DB_HOST      string `mapstructure:"DB_HOST"`
	DB_PORT      string `mapstructure:"DB_PORT"`
	DB_USER      string `mapstructure:"DB_USER"`
	DB_DRIVER    string `mapstructure:"DB_DRIVER"`
	DB_PASSWORD  string `mapstructure:"DB_PASSWORD"`
	DB_NAME      string `mapstructure:"DB_NAME"`
	API_PORT     string `mapstructure:"API_PORT"`
	JWT_SECRET   string `mapstructure:"JWT_SECRET"`
	JWT_EXPIREIN int    `mapstructure:"JWT_EXPIREIN"`
	TOKEN_AUTH   *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("my_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TOKEN_AUTH = jwtauth.New("HS256", []byte(cfg.JWT_SECRET), nil)

	return cfg, err
}
