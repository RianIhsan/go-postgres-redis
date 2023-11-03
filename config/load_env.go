package config

import "github.com/spf13/viper"

type Config struct {
	// Postgres setup
	DbHost     string `mapstructure:"DBHOST"`
	DbUsername string `mapstructure:"DBUSER"`
	DbPassword string `mapstructure:"DBPASS"`
	DbName     string `mapstructure:"DBNAME"`
	DbPort     string `mapstructure:"DBPORT"`
	DbSSL      string `mapstructure:"DBSSL"`
	TimeZone   string `mapstructure:"TIMEZONE"`

	//Redis setup
	RedisUrl string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
