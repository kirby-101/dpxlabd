package dpxlabd

/*
  - build decision controller

* src+port bedingungen aus config holen
* relevanter commit (branch art?, letzder build? wegen cooldown)
* was hat sich geändert / relevant? /
* branch?
* letzder build aufgrund cooldown / grade ein build schon aktiv?
* commits auf pkg makefile versionierung fokussieren?
* build-dep: src -> port/port-update ; sodass nie inkonsistent ist

bei radicle git fragen
*/

/*
 * EngineController: make build decisions
 */
type SchedEngineController struct {
	// []Trigger
	// []Job
	//Queue
}

/*
	// buildWG tracks in-flight build goroutines so shutdown can let their
	// deferred cleanup run. Without it main returns as soon as repoWG drains,
	// the process exits, and every build's chroot copy and build tree is left
	// behind with nothing logged, because a goroutine that dies with the
	// process cannot report it.
	buildWG   *sync.WaitGroup
*/

func (c *SchedEngineController) EnqueueRunJob() {}

/*
// trigger/job queue?


   Mutex Lock: nur 1 job aufeinmal
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


type SchedQueue struct {
	jobLock sync.Mutex
}

func NewSchedQueue() *SchedQueue {}

//func (q *SchedQueue)


type BuilderQueue struct {
	elems []Job
	seen  map[Job]struct{}
}

func NewBuilderQueue() *BuilderQueue {
	return &BuilderQueue{seen: make(map[Job]struct{})}
}

// to poll?
func (q *BuilderQueue) Len() int {
	return len(q.elems)
}

func (q *BuilderQueue) Enqueue(j Job) {
	if _, exists := q.seen[j]; !exists {
		// enqueue
		q.elems = append(q.elems, j)
		q.seen[j] = struct{}{}
	}
}

func (q *BuilderQueue) Dequeue() Job {

	// return front Job & Dequeue item
	return nil
}
*/
