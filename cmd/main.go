package main

import (
	"context"
	"github.com/glebnaz/notion-recurring-tasks/internal/server"
	"github.com/glebnaz/notion-recurring-tasks/internal/service"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// Release is a semantic version of current build
	Release = "unset"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
}

type config struct {
	NotionToken  string `envconfig:"NOTION_TOKEN"`
	PathToConfig string `envconfig:"PATH_TO_CONFIG" default:"config.json"`
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		ForceColors:   true,
		DisableQuote:  true,
	})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

func main() {
	log.Infof("Starting App. Release: %s Commit: %s  BuildTime: %s", Release, Commit, BuildTime)
	s := server.NewServer()

	ctx := context.Background()

	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Errorf("err read config: %s", err)
		panic(err)
	}

	taskConfig, err := service.NewConfigFromFile(cfg.PathToConfig)
	if err != nil {
		log.Errorf("err read config: %s", err)
		panic(err)
	}

	ctrl := service.NewRecurringTaskController(taskConfig, cfg.NotionToken)

	err = ctrl.RegisterConfig(ctx)
	if err != nil {
		log.Errorf("err read config: %s", err)
		panic(err)
	}

	ctrl.Start(ctx)

	s.Run()
}
