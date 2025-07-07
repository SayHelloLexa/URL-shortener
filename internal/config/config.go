package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Domain string `mapstructure:"domain"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") 
	viper.SetConfigType("yaml")   
	viper.AddConfigPath("internal/config") 

	// Установка значений по умолчанию
	viper.SetDefault("domain", "localhost:8080")

	// Чтение конфигурации
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Файл конфигурации не найден, используются значения по умолчанию")
		} else {
			return nil, fmt.Errorf("ошибка чтения конфигурации: %v", err)
		}
	}

	// Автоматическое чтение переменных окружения
	viper.AutomaticEnv()

	// Декодирование конфигурации в структуру
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка декодирования конфигурации: %v", err)
	}

	return &cfg, nil
}
