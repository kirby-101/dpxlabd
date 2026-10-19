package dpxlabd

import (
	"sync"

	_ "github.com/hetznercloud/hcloud-go/v2/hcloud"
)

/*
* x86-64-v2
* x86-64-v3
 */

/*
pkg.dpxlab.de/osreldate_portscommit
http://pkg.dpxlab.de/1501502_ports-git-tag/
*/

/*
(head_src, head_ports) != (last_built_src, last_built_ports)
180h cooldown
security/vuxml/vuln.xml prüfen
bei security advisory, direkt bauen & cooldown ignorieren

es sind nd 2 jobs sondern eine pipeline
zuerst src & ggf. ports
buildkernel buildwold ; ports neu gegen neue src

Was bei failed builds?
warten auf neuen commit mit fix?
Wenn src fails, nicht ports bauen
*/
type SchedEngine struct {
	sync.Mutex // nur 1 job aufeinmal
}

// Start & Ende benachrichtigen; & kosten
func (e *SchedEngine) RunJob(t Trigger) {
	e.Lock()
	defer e.Unlock()

	// middleware
	/*
	 * welche OS images bzw src haben wir?
	 * welcher job?
	 * welcher commit (für checkout)
	 * compute Engine zum bauen
	 * repo in bucket storage
	 */
}
