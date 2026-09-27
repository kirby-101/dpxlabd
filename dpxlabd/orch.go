package dpxlabd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kirby-101/dpxlab/config"
	"github.com/kirby-101/dpxlab/db"
	"github.com/kirby-101/logging-go"
)

// ? store;sched;ddn
type Orchestrator struct {
	www struct {
		routes map[string]RouteHandler
		mux    *http.ServeMux
		errR   RouteHandler
	}

	cfg    *config.Config
	db     *db.Database
	Logger *logging.Logger

	signals chan os.Signal
}

func New(
	cfg *config.Config,
	db *db.Database,
	logger *logging.Logger,
	httpRoutes map[string]RouteHandler,
	errRoute RouteHandler,
) *Orchestrator {
	var orch = &Orchestrator{
		cfg:     cfg,
		db:      db,
		Logger:  logger,
		signals: make(chan os.Signal, 1),
	}

	orch.www.errR = errRoute
	orch.www.routes = httpRoutes
	orch.www.mux = http.NewServeMux()

	return orch
}

func (o *Orchestrator) Start() {
	signal.Notify(
		o.signals,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	o.run()
}

func (o *Orchestrator) Stop() {
	o.handlePanic()
	os.Exit(0)
}

func (o *Orchestrator) run() {
	defer o.Stop()

	// http listener
	var c_err = make(chan error, 1)
	go func() {
		c_err <- o.StartWWW()
	}()

	for {
		select {
		case err := <-c_err:
			// fatal error
			o.Logger.Error("HTTP Listener", fmt.Sprintf("%#v", err))
			o.Stop()

		case <-o.signals:
			o.Logger.Info("catched SIGINT/SIGTERM")
			o.Stop()
		}
	}
}

func (o *Orchestrator) handlePanic() {
	if r := recover(); r != nil {
		o.Logger.Fatal("PANIC", fmt.Sprintf("%#v", r))
	}
}
