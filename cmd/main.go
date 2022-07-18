package main

import (
	"github.com/glebnaz/notion-recurring-tasks/internal/server"
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
	log.SetFormatter(&log.JSONFormatter{})
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

	s.Run()
}
