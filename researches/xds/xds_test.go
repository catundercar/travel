package xds

import (
	"github.com/heiyeluren/xds/xmap"
	"github.com/heiyeluren/xmm"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkCHM_Concurrent_Get(b *testing.B) {
	f := &xmm.Factory{}
	mm, err := f.CreateMemory(0.75)
	if err != nil {
		b.Fatal(err)
	}
	chm, err := xmap.NewConcurrentRawHashMap(mm, 16, 2, 8)
	if err != nil {
		b.Fatal(err)
	}
	keys := make([]string, 8000000)
	for i := 0; i < 8000000; i++ {
		keys[i] = strconv.Itoa(rand.Int() % 800000)
	}
	for _, key := range keys {
		if err := chm.Put([]byte(key), []byte(key)); err != nil {
			b.Error(err)
		}
	}
	length := len(keys)
	
	b.SetParallelism(100)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := keys[rand.Int()%length]
			if _, exist, err := chm.Get([]byte(key)); err != nil || !exist {
				b.Error(err)
			}
		}
	})
}

func BenchmarkSyncMap_Get(b *testing.B) {
	b.ReportAllocs()
	synMap := sync.Map{}
	keys := make([]string, 8000000)
	for i := 0; i < 8000000; i++ {
		keys[i] = strconv.Itoa(rand.Int() % 800000)
	}
	for _, key := range keys {
		synMap.Store(key, key)
	}
	length := len(keys)

	b.SetParallelism(100)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := keys[rand.Int()%length]
			if _, ok := synMap.Load(key); !ok {
				b.Error("not found")
			}
		}
	})
}
