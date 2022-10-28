package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	f := sync.Map{}

	keys := make([]string, 8000000)
	for i := 0; i < 8000000; i++ {
		keys[i] = strconv.Itoa(rand.Int() % 800000)
	}
	for _, key := range keys {
		f.Store(key, key+"xx")
	}

	http.HandleFunc("/sync_map/get", func(w http.ResponseWriter, r *http.Request) {
		key := keys[rand.Int()%8000000]
		v, exist := f.Load(key)
		if !exist {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf(`{"error": "not found"}`)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"%s": "%s"}`, key, v)))
	})
	if err := http.ListenAndServe(":8089", nil); err != nil {
		panic(err)
	}
}
