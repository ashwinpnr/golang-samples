package config

import (
	"log"

	"github.com/spf13/viper"
)

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

var CurrentConfig Configurations

func SetConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&CurrentConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
func GetConfig() Configurations {
	return CurrentConfig
}

func init() {
	SetConfig()
}
