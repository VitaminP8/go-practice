package additional_synchronization_primitives

import "sync"

type Worker struct {
	ready bool
	mu    sync.RWMutex
}

func (w *Worker) setReady() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.ready = true
}
func (w *Worker) CheckReady() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.ready
}
