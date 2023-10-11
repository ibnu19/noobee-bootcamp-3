package sesi2concurrency

import (
	"fmt"
	"sync"
	"testing"
)

type M map[string]string

func PrintData(key, value string) {
	fmt.Printf("%s:%s\n", key, value)
}

func TestGoroutine(t *testing.T) {
	arg := map[string]string{
		"Name":    "NooBee",
		"Class":   "Backend Intermediate",
		"Address": "Jakarta",
	}

	wg := sync.WaitGroup{}

	for Key, value := range arg {
		wg.Add(1)
		go func(key, value string) {
			PrintData(key, value)
			wg.Done()
		}(Key, value)
	}

	wg.Wait()
}
