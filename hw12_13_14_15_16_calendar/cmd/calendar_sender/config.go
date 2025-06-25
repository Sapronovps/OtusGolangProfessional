package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	RabbitMQ RabbitMQConf
	Logger   LoggerConf
}

type RabbitMQConf struct {
	User         string
	Password     string
	Host         string
	Port         string
	QueueName    string
	ExchangeName string
	RoutingKey   string
}

type LoggerConf struct {
	File  string
	Level string
}

func NewConfig(configPath string) Config {
	viper.SetConfigName("config-sender")
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

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения конфигурации: %w", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("ошибка декодирования конфигурации в структуру: %w", err))
	}

	return config
}
