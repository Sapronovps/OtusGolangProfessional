package main

import (
	"fmt"
	"github.com/spf13/viper"
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
	Host string
	Port int
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
	DbName   string
	InMemory bool
}

func NewConfig(configPath string) Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.SetEnvPrefix("db")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения конфигурации: %w", err))
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("decode into struct: %w", err))
	}

	return c
}
