package main

import (
	"github.com/glebnaz/notion-recurring-tasks/internal/server"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
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

func main() {
	log.Infof("Starting...")
	go print()

	srv := server.NewServer()

	srv.AddChecker(server.NewDefaultChecker("simple", func() error {
		return nil
	}))

	srv.Run()

}

func print() {
	for {
		time.Sleep(time.Second)
		log.Infof("Printing...")
	}
}
