package dpxlabd

import (
	"sync"

	_ "cloud.google.com/go/compute/apiv1"
	_ "cloud.google.com/go/compute/apiv1/computepb"
	_ "google.golang.org/protobuf/proto"
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

bei ACPI G2 Soft Off / SIGTERM / SIGINT bzw shutdown, standy packages auf r2 pushen (120sec)
cache und packages sofort hochladen, nicht warten
*/
type SchedEngine struct {
	sync.Mutex // nur 1 job aufeinmal
}

// Start & Ende benachrichtigen; & kosten
func (e *SchedEngine) RunJob(t SchedTrigger) {
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

/*
   * Spot VM!
   * Prod: c3d-highcpu-16 0,15€/h => ~5$ bis 12$
   * Debug-VM: e2-medium ~0,02$/h / e2-micro gratis in us-west1,us-central1,us-east1
   * Kosten?: ~140h bis 175h für Kaltstart
   * ccache, ALLOW_MAKE_JOBS, tmpfs, gute PARALLEL_JOBS
   * NETWORK IST TEUER

gcloud compute instances create my-app-instance \
--image-family=debian-10 \
--image-project=debian-cloud \
--machine-type=g1-small \
--scopes userinfo-email,cloud-platform \
--metadata-from-file startup-script=startup-script.sh \
--metadata app-location="gs://YOUR_BUCKET_NAME/app.tar.gz" \
--zone YOUR_ZONE \
--tags http-server

gcloud compute instances get-serial-port-output my-app-instance --zone YOUR_ZONE

gcloud compute firewall-rules create default-allow-http-80 \
    --allow tcp:80 \
    --source-ranges 0.0.0.0/0 \
    --target-tags http-server \
    --description "Allow port 80 access to http-server"

*/

/*
Google Cloud Compute Engine (Spot oder On-Demand)
Am besten passend, weil du nur stundenweise zahlst und die Maschine nach dem Build einfach stoppen oder löschen kannst.
Für FreeBSD-Builds nimm mindestens 4 vCPU + 8–16 GB RAM (e2-standard-4 oder e2-highcpu-8).
Spot-VMs sind deutlich günstiger und für Builds meist okay (falls sie unterbrochen werden, einfach neu starten).
Persistent Disk kannst du klein halten oder nach dem Build die Packages nach Backblaze B2 / R2 / Wasabi hochladen und die Disk löschen.
*/
