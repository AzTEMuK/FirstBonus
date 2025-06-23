package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    Server struct {
        Host string
        Port int
    }
}

var cfg *Config

func Load(path string) {
    viper.SetConfigFile(path)
    viper.SetConfigType("yaml")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Ошибка чтения конфига: %v", err)
    }

    cfg = &Config{}
    if err := viper.Unmarshal(cfg); err != nil {
        log.Fatalf("Ошибка разбора конфига: %v", err)
    }
}

func Get() *Config {
    if cfg == nil {
        log.Fatal("Конфигурация не загружена. Вызови config.Load() перед использованием.")
    }
    return cfg
}