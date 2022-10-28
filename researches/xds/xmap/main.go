package main

import (
	"fmt"
	"github.com/heiyeluren/xds/xmap"
	"github.com/heiyeluren/xmm"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	f := &xmm.Factory{}
	mm, err := f.CreateMemory(0.75)
	if err != nil {
		panic(err)
	}

	chm, err := xmap.NewConcurrentRawHashMap(mm, 16, 2, 8)
	if err != nil {
		panic(err)
	}
	keys := make([]string, 8000000)
	for i := 0; i < 8000000; i++ {
		keys[i] = strconv.Itoa(rand.Int() % 800000)
	}
	for _, key := range keys {
		if err := chm.Put([]byte(key), []byte(key+"xx")); err != nil {
			panic(err)
		}
	}

	http.HandleFunc("/xmap/get", func(w http.ResponseWriter, r *http.Request) {
		key := keys[rand.Int()%8000000]
		v, exist, err := chm.Get([]byte(key))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"error": "%v"}`, err.Error())))
			return
		} else if !exist {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf(`{"error": "not found"}`)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"%s": "%s"}`, key, v)))
	})
	if err = http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
