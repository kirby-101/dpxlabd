package dpxlabd

import "context"

type Sched struct {
	//? storage, engine, controller
}

func (srv *Server) StartSched(errChannel chan error) {
	if !srv.Config.Sched.Enabled {
		srv.Logger.Info("Sched disabled")
		return
	}

	// start controller / engine
}

func (srv *Server) StatusSched() {}

func (srv *Server) StopSched(ctx context.Context) {
	// is running?
	// stop engine
}
