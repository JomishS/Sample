package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	/*this provides metadata about how this field should be mapped when unmarshalling data from a configuration file like yaml or json.
	here it specifies that when unmarshalling the data ,the field should be mapped from the key database in the configuration source*/
}

var config Configuration

var ViperConfig *viper.Viper

func Loadconfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	ViperConfig = viper.GetViper()

	err = ViperConfig.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)

		return
	}
}

func GetConfig() Configuration {
	return config
}
