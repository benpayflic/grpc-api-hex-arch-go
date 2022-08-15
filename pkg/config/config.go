package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongoURI            string `mapstructure:"MONGO_URI"`
	MongoDBame          string `mapstructure:"MONGO_DB_NAME"`
	MongoCollectionName string `mapstructure:"MONGO_COLLECTION_NAME"`
}

// Load configurations from path and unmarshal to Config
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
