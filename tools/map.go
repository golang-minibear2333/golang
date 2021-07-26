package tools

import "sync"

type M interface {
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{})
}

type Initializer func() interface{}

type GuardedM struct {
	sync.RWMutex
	M M
	i Initializer
}

func (gm *GuardedM) Init(m M, initer Initializer) *GuardedM {
	gm.M = m
	gm.i = initer
	return gm
}

func (gm *GuardedM) Get(k interface{}) interface{} {
	gm.RLock()
	m, ok := gm.M.Get(k)
	gm.RUnlock()

	if ok {
		return m
	}

	gm.Lock()
	if _, ok := gm.M.Get(k); !ok {
		m = gm.i()
		gm.M.Set(k, m)
	}
	gm.Unlock()

	return m
}

func (gm *GuardedM) Set(key, value interface{}) {
	gm.Lock()
	gm.M.Set(key, value)
	gm.Unlock()
}
