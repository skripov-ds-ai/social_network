package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// type Config struct {
//	DBConfig
//	ServerPort string `mapstructure:"PORT"`
//}

type Config struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	DB       string `mapstructure:"POSTGRES_DB"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`

	ServerPort string `mapstructure:"PORT"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	// viper.SetConfigName("app")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	err := viper.Unmarshal(&config)
	return config, err
}

func (c *Config) URI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DB)
}
