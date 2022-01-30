package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	appPort    = "APP_PORT"
	appHost    = "APP_HOST"
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUserName = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
)

type Config struct {
	AppPort    string
	AppHost    string
	DbHost     string
	DbPort     string
	DbUserName string
	DbPassword string
	DbName     string
}

func (c *Config) Valid() bool {
	return c.AppPort != "" && c.AppHost != "" && c.DbHost != "" && c.DbPort != "" && c.DbUserName != "" && c.DbPassword != "" && c.DbName != ""
}

func GetConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := &Config{
		AppPort:    viper.GetString(appPort),
		AppHost:    viper.GetString(appHost),
		DbHost:     viper.GetString(dbHost),
		DbPort:     viper.GetString(dbPort),
		DbUserName: viper.GetString(dbUserName),
		DbPassword: viper.GetString(dbPassword),
		DbName:     viper.GetString(dbName),
	}
	if !config.Valid() {
		return nil, errors.New("config is invalid")
	}
	return config, nil
}
