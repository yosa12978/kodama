package config

import (
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		LogLevel string `yaml:"log_level"`
		Addr     string `yaml:"addr"`
	} `yaml:"app"`
	Postgres struct {
		Database string `yaml:"database"`
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"postgres"`
}

func DefaultConfig() Config {
	return Config{
		App: struct {
			LogLevel string "yaml:\"log_level\""
			Addr     string "yaml:\"addr\""
		}{
			LogLevel: "INFO",
			Addr:     "0.0.0.0:80",
		},
		Postgres: struct {
			Database string "yaml:\"database\""
			Addr     string "yaml:\"addr\""
			User     string "yaml:\"user\""
			Password string "yaml:\"password\""
		}{
			Database: "kodama_db",
			Addr:     "localhost:5432",
			User:     "postgres",
			Password: "postgres",
		},
	}
}

func ReadConfig(filename string) (Config, error) {
	conf := DefaultConfig()
	file, err := os.Open(filename)
	if err != nil {
		return conf, err
	}
	defer file.Close()
	err = yaml.NewDecoder(file).Decode(&conf)
	if errors.Is(err, io.EOF) {
		return conf, nil
	}
	return conf, err
}
