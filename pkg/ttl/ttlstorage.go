package ttl

import (
	"sync"
	"time"
)

type TTLData struct {
	Val int
	Exp time.Time
}

type TTLMap struct {
	M     map[string]TTLData
	TTL   time.Duration
	mutex sync.RWMutex
}

func (m *TTLMap) Init(d time.Duration) *TTLMap {
	m0 := &TTLMap{
		M:   make(map[string]TTLData),
		TTL: d,
	}
	go m0.mapCleaner(d)
	return m0
}

func (m *TTLMap) mapCleaner(d time.Duration) {
	for {
		<-time.After(d)
		m.mutex.Lock()
		for k, v := range m.M {
			if time.Now().After(v.Exp) {
				delete(m.M, k)
			}
		}
		m.mutex.Unlock()
	}
}

func (m *TTLMap) Inc(key string) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	val, ok := m.M[key]
	switch ok {
	case true:
		if time.Now().After(val.Exp) {
			delete(m.M, key)
		} else {
			val.Val += 1
		}
	case false:
		val.Val = 1
		val.Exp = time.Now().Add(m.TTL)
	}
	m.M[key] = val
	return val.Val
}
