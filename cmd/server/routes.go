package main

import (
	"github.com/kirby-101/dpxlabd"
	"github.com/kirby-101/dpxlabd/routes/adm"
	"github.com/kirby-101/dpxlabd/routes/index"
	"github.com/kirby-101/dpxlabd/routes/markdown"
	"github.com/kirby-101/dpxlabd/routes/robots"
	"github.com/kirby-101/dpxlabd/routes/static"
)

func registerAllSubroutes(srv *dpxlabd.Server) {
	var (
		_ = index.New(srv).Register(srv.WWW.Router)
		_ = robots.New(srv).Register(srv.WWW.Router)
		_ = markdown.New(srv).Register(srv.WWW.Router)
		_ = static.New(srv).Register(srv.WWW.Router)
		_ = adm.New(srv).Register(srv.WWW.Router)
	)
}
