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

func CanopusNodes() []uint64 {
	var nodes []uint64
	for i := 1; i < 6; i++ {
		nodes = append(nodes, uint64(i))
	}
	return nodes
}
