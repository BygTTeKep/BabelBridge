package config

import (
	"strings"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
}

type AppConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DeepLConfig struct {
	AuthKey string `mapstructure:"authKey"`
	BaseURL string `mapstructure:"baseUrl"`
}

type Config struct {
	AppCfg   AppConfig      `mapstructure:"app"`
	DBCfg    DatabaseConfig `mapstructure:"database"`
	DeepLCfg DeepLConfig    `mapstructure:"deepL"`
}

func LoadConfig(path string) (conf *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}
	err = viper.Unmarshal(&conf)
	return
}
