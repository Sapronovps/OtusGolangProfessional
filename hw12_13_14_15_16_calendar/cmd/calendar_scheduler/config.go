package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	RabbitMQ RabbitMQConf
	DB       DBConf
	Logger   LoggerConf
}

type RabbitMQConf struct {
	User         string
	Password     string
	Host         string
	Port         string
	ExchangeName string
	RoutingKey   string
	ScanInterval int64
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

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Задаем переменные окружения вручную, чтобы указать маппинг
	viper.BindEnv("rabbitmq.user")
	viper.BindEnv("rabbitmq.password")
	viper.BindEnv("rabbitmq.host")
	viper.BindEnv("rabbitmq.port")
	viper.BindEnv("rabbitmq.exchangename")
	viper.BindEnv("rabbitmq.routingkey")
	viper.BindEnv("rabbitmq.scaninterval")

	viper.BindEnv("logger.file")
	viper.BindEnv("logger.level")

	viper.BindEnv("db.host")
	viper.BindEnv("db.port")
	viper.BindEnv("db.username")
	viper.BindEnv("db.password")
	viper.BindEnv("db.dbname")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения конфигурации: %w", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("ошибка декодирования конфигурации в структуру: %w", err))
	}

	return config
}
