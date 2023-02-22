package main

import (
	"fmt"
	"go-storage/pkg/ttl"
	"time"
)

func main() {
	var storage ttl.TTLMap

	ttl := time.Second * 5
	s := storage.Init(ttl)

	fmt.Println(s.Inc("key1"))
	fmt.Println(s.Inc("key2"))
	fmt.Println(s.Inc("key1"))
	time.Sleep(ttl)
	fmt.Println(s.Inc("key1"))
}
