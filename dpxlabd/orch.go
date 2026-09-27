package dpxlabd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
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
	sync.Mutex
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

func (o *Orchestrator) Start() error {
	o.Lock()

	signal.Notify(
		o.signals,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	// 	o.StartWWW()

	return o.run()
}

func (o *Orchestrator) Stop() {
	defer o.handlePanic()
	o.Lock()

	os.Exit(0)
}

func (o *Orchestrator) run() error {
	defer o.Stop()

	for {
		select {
		//case X:

		case <-o.signals:
			o.Logger.Info("catched SIGINT/SIGTERM")
			return nil

		}
	}
}

func (o *Orchestrator) handlePanic() {
	if r := recover(); r != nil {
		o.Lock()
		o.Logger.Fatal("PANIC", fmt.Sprintf("%#v", r))
	}
}
