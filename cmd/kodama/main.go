package main

import (
	"errors"
	"flag"
	"log/slog"
	"os"

	"github.com/yosa12978/kodama/internal/app"
	"github.com/yosa12978/kodama/internal/config"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.yml", "kodama config file path")
}

func main() {
	flag.Parse()

	appConfig, err := config.ReadConfig(configFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Warn("Config file doesn't exist. Switching to defaults.", "file", configFile)
		} else {
			panic(err)
		}
	}

	kodama := app.NewFromConfig(appConfig)
	if err := kodama.Run(); err != nil {
		panic(err)
	}
}
