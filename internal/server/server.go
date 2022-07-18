package server

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type engineCfg struct {
	DebugPort       string        `json:"debug_port" envconfig:"DEBUG_PORT" default:":8084"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout" envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

//todo нужно добавить фоновые задачи с отключением, и абстракцию которая сможет проверять лайф
type Server struct {
	*DebugServer

	shutdownTimeout time.Duration
}

func (s *Server) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	log.Infof("Run Server")
	go func() {
		err := s.DebugServer.Run(ctx)
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("Error Run Debug Server: %s", err)
		}
	}()
	s.SetReady(true)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Infof("Start shutdown server")
	s.Shutdown(ctx)
	cancel()
	log.Infof("Shutdown server")
}

func NewServer() Server {
	var cfg engineCfg

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Errorf("err read engine config: %s", err)
		panic(err)
	}

	debug := NewDebugServer(cfg.DebugPort)

	s := Server{DebugServer: debug, shutdownTimeout: cfg.ShutdownTimeout}
	return s
}
