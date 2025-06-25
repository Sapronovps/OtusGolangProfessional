package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Server ServerConf
	Logger LoggerConf
	DB     DBConf
}

type ServerConf struct {
	Host        string
	Port        int
	AddressGrpc string
	IsHTTP      bool
}

type LoggerConf struct {
	File  string
	Level string
}

type DBConf struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	InMemory bool
}

func NewConfig(configPath string) Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Задаем переменные окружения вручную, чтобы указать маппинг
	viper.BindEnv("server.host")
	viper.BindEnv("server.port")
	viper.BindEnv("server.addressgrpc")
	viper.BindEnv("server.ishttp")

	viper.BindEnv("logger.file")
	viper.BindEnv("logger.level")

	viper.BindEnv("db.host")
	viper.BindEnv("db.port")
	viper.BindEnv("db.username")
	viper.BindEnv("db.password")
	viper.BindEnv("db.dbname")
	viper.BindEnv("db.inmemory")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения конфигурации: %w", err))
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("decode into struct: %w", err))
	}

	return c
}
