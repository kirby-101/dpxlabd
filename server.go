package dpxlabd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	_config "github.com/kirby-101/dpxlabd/config"
	logging "github.com/kirby-101/logging-go"
)

type Server struct {
	WWW    *WWW
	Sched  *Sched
	Config *_config.Config
	Logger *logging.Logger
}

func New(config *_config.Config, logger *logging.Logger) *Server {
	var (
		router = mux.NewRouter()
		srv    = &Server{
			WWW: &WWW{
				Router: router,
				srv: &http.Server{
					Addr:         config.WWW.Address,
					Handler:      router,
					WriteTimeout: 15 * time.Second,
					ReadTimeout:  15 * time.Second,
					IdleTimeout:  time.Second * 60,
				},
			},
			Config: config,
			Logger: logger,
		}
	)

	return srv
}

func (srv *Server) Start() {
	srv.Logger.Info("Starting...")

	// Signals
	var sigs = make(chan os.Signal, 1)
	signal.Notify(
		sigs,
		syscall.SIGINFO,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	var errChannel = make(chan error, 1)

	// www server
	go srv.StartWWW(errChannel)

	// sched
	go srv.StartSched(errChannel)

	srv.run(sigs, errChannel)
}

func (srv *Server) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5, //! DEADLINE
	)
	defer cancel()

	// www server
	go srv.StopWWW(ctx)

	// sched
	go srv.StopSched(ctx)

	srv.Logger.Info("waiting 15s shutdown timeout")
	<-ctx.Done()

	srv.handlePanic()
	os.Exit(0)
}

func (srv *Server) Status() {}

func (srv *Server) run(sigs chan os.Signal, errChannel chan error) {
	defer srv.Stop()

	for {
		select {
		case err := <-errChannel:
			// fatal
			srv.Logger.Error(err.Error())
			srv.Stop()

		case sig := <-sigs:
			srv.handleSignal(sig)
		}
	}
}

func (srv *Server) handlePanic() {
	if r := recover(); r != nil {
		srv.Logger.Fatal("PANIC", fmt.Sprintf("%#v", r))
	}
}

func (srv *Server) handleSignal(sig os.Signal) {
	switch sig {
	case syscall.SIGINFO:
		srv.Status()

	case syscall.SIGINT, syscall.SIGTERM:
		srv.Logger.Info("catched SIGINT/SIGTERM")
		srv.Stop()
	}
}
