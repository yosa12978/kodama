package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/yosa12978/kodama/internal/app"
	"github.com/yosa12978/kodama/internal/config"
)

var configFile string

const (
	noConfigMsg = "\x1b[1;33mConfig file \"%s\" doesn't exist. Using default config.\x1b[00m\n"
)

func init() {
	flag.StringVar(&configFile, "config", "/etc/kodama/kodama.yml", "kodama config file path")
}

func main() {
	flag.Parse()

	appConfig, err := config.ReadConfig(configFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintf(os.Stderr, noConfigMsg, configFile)
		} else {
			panic(err)
		}
	}

	kodama := app.NewFromConfig(appConfig)
	if err := kodama.Run(); err != nil {
		panic(err)
	}
}
