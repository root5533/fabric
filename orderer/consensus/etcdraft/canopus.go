package etcdraft

import "sync"

type WriteLock struct {
	writing bool
	mux     sync.Mutex
}

func (lock *WriteLock) getWriteVal() bool {
	lock.mux.Lock()
	defer lock.mux.Unlock()
	return lock.writing
}

func (lock *WriteLock) setWriteVal(val bool) {
	lock.mux.Lock()
	lock.writing = val
	lock.mux.Unlock()
}

func CanopusNodes() []int {
	return []int{1, 2, 3, 4, 5}
}
