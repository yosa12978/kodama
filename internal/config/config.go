package config

import (
	"errors"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		LogLevel    string `yaml:"log_level"`
		Addr        string `yaml:"addr"`
	} `yaml:"app"`
	Postgres struct {
		Database     string `yaml:"database"`
		Addr         string `yaml:"addr"`
		User         string `yaml:"user"`
		UserFile     string `yaml:"user_file"`
		Password     string `yaml:"password"`
		PasswordFile string `yaml:"password_file"`
	} `yaml:"postgres"`
}

func DefaultConfig() Config {
	return Config{
		App: struct {
			Name        string "yaml:\"name\""
			Description string "yaml:\"description\""
			LogLevel    string "yaml:\"log_level\""
			Addr        string "yaml:\"addr\""
		}{
			Name:        "Kodama",
			Description: "",
			LogLevel:    "INFO",
			Addr:        "0.0.0.0:80",
		},
		Postgres: struct {
			Database     string "yaml:\"database\""
			Addr         string "yaml:\"addr\""
			User         string "yaml:\"user\""
			UserFile     string "yaml:\"user_file\""
			Password     string "yaml:\"password\""
			PasswordFile string "yaml:\"password_file\""
		}{
			Database:     "kodama_db",
			Addr:         "localhost:5432",
			User:         "postgres",
			UserFile:     "",
			Password:     "postgres",
			PasswordFile: "",
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
	if err := parseFileConfigs(&conf); err != nil {
		return conf, err
	}
	return conf, err
}

func parseFileConfigs(c *Config) error {

	if c.Postgres.UserFile != "" {
		userFile, err := os.ReadFile(c.Postgres.UserFile)
		if err != nil {
			return err
		}
		c.Postgres.User = strings.TrimSpace(string(userFile))
	}

	if c.Postgres.PasswordFile != "" {
		passwordFile, err := os.ReadFile(c.Postgres.PasswordFile)
		if err != nil {
			return err
		}
		c.Postgres.Password = strings.TrimSpace(string(passwordFile))
	}

	return nil
}
