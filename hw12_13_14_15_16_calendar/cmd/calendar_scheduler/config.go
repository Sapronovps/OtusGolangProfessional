package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	RabbitMQ RabbitMQConf
	DB       DBConf
	Logger   LoggerConf
}

type RabbitMQConf struct {
	User     string
	Password string
	Host     string
	Port     string
}

type DBConf struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type LoggerConf struct {
	File  string
	Level string
}

func NewConfig(configPath string) Config {
	viper.SetConfigName("config-scheduler")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.SetEnvPrefix("scheduler")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения конфигурации: %w", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("ошибка декодирования конфигурации в структуру: %w", err))
	}

	return config
}
