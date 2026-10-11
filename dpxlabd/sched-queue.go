package dpxlabd

import "sync"

// trigger/job queue?

/*
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
*/

type SchedQueue struct {
	jobLock sync.Mutex
}

func NewSchedQueue() *SchedQueue {}

//func (q *SchedQueue)

/*
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
