package dpxlabd

import _ "github.com/hetznercloud/hcloud-go/v2/hcloud"

/* https://pkg.go.dev/github.com/hetznercloud/hcloud-go/v2/hcloud */

/*

pkg.dpxlab.de/osreldate_portscommit
http://pkg.dpxlab.de/1501502/ports-git-tag
*/

func (o *Orchestrator) StartSched() error {
	// launch triggers
}

// abort current job
func (o *Orchestrator) StopSched() error {}

// keybind/signal => print status
func (o *Orchestrator) Status() {}

/*
poudriere bulk -v -j hbsd-14-stable -p default -O hyprland_overlay -f /usr/local/etc/poudriere.d/packages-default

poudriere sets für desktop und laptop
*/
