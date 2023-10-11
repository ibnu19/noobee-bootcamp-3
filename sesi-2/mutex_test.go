package sesi2concurrency

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

// Goroutine without
func TestGoroutineWithoutMutex(t *testing.T) {
	var A = 1000
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			A++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Total A:", A)
}

// Goroutine with mutex
type counter struct {
	count int
	sync.Mutex
}

func (c *counter) Add() {
	c.Lock()
	c.count++
	c.Unlock()
}

func TestCounterWithMutex(t *testing.T) {
	c := counter{
		count: 0,
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Add()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("Counter:", c.count)
}
