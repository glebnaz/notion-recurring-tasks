package server

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

type DebugServer struct {
	PORT string

	engine *echo.Echo
	m      sync.Mutex

	ready    bool
	checkers []Checker

	shutdownTimeout time.Duration
}

func NewDebugServer(port string) *DebugServer {
	e := echo.New()
	e.Debug = false
	e.HideBanner = true
	e.HidePort = true

	debug := &DebugServer{
		PORT:            port,
		engine:          e,
		ready:           false,
		checkers:        make([]Checker, 0),
		shutdownTimeout: time.Second * 30,
	}

	e.GET("/ready", debug.Ready)
	e.GET("/live", debug.Live)

	return debug
}

func (d *DebugServer) SetReady(ready bool) {
	d.m.Lock()
	defer d.m.Unlock()
	d.ready = ready
	if ready {
		log.Infof("Server is ready")
	} else {
		log.Infof("Server is not ready")
	}
}

func (d *DebugServer) AddChecker(checker Checker) {
	log.Debugf("Adding checker %s", checker.Name())
	d.m.Lock()
	defer d.m.Unlock()
	d.checkers = append(d.checkers, checker)
}

func (d *DebugServer) AddCheckers(checkers []Checker) {
	for _, checker := range checkers {
		d.AddChecker(checker)
	}
}

func (d *DebugServer) Ready(c echo.Context) error {
	d.m.Lock()
	defer d.m.Unlock()
	return c.JSON(200, map[string]interface{}{
		"ready": d.ready,
	})
}

func (d *DebugServer) Live(c echo.Context) error {
	d.m.Lock()
	defer d.m.Unlock()
	for i := range d.checkers {
		if err := d.checkers[i].Check(); err != nil {
			log.Errorf("Checker %s failed: %s", d.checkers[i].Name(), err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}
	return c.String(http.StatusOK, "OK")
}

type Checker interface {
	Check() error
	Name() string
}

type DefaultChecker struct {
	CheckFunc func() error `json:"check"`
	NameCheck string       `json:"name"`
}

func (c *DefaultChecker) Check() error {
	return c.CheckFunc()
}

func (c *DefaultChecker) Name() string {
	return c.NameCheck
}

func NewDefaultChecker(name string, checkFunc func() error) *DefaultChecker {
	return &DefaultChecker{
		CheckFunc: checkFunc,
		NameCheck: name,
	}
}
