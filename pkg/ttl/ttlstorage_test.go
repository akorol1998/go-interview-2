package ttl

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func incrementVal(m *TTLMap, idx int) {
	m.Inc(fmt.Sprintf("key%d", idx))
}

func BenchmarkSingleRoutine(b *testing.B) {
	var storage TTLMap
	ttl := time.Second * 5

	s := storage.Init(ttl)
	for i := 0; i < b.N; i++ {
		incrementVal(s, i)
	}
}

func BenchmarkMultipleRoutines(b *testing.B) {
	var wg sync.WaitGroup
	var storage TTLMap

	ttl := time.Second * 5
	wg.Add(b.N)
	s := storage.Init(ttl)
	for i := 0; i < b.N; i++ {
		go func(idx int) {
			incrementVal(s, idx)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
