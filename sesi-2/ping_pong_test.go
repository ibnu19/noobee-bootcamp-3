package sesi2concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type PlayerX struct {
	Name string
	Hit  int
}

const BreakPoints int = 11

func TestPingPongApp(t *testing.T) {
	ping := make(chan *PlayerX)
	pong := make(chan *PlayerX)
	isDone := make(chan *PlayerX)
	numChan := make(chan int)

	A := PlayerX{Name: "Player A", Hit: 0}
	B := PlayerX{Name: "Player B", Hit: 0}

	// Player A
	go pinger(ping, pong, isDone, &A, numChan)
	// Player B
	go ponger(ping, pong, isDone, &B, numChan)

	ping <- new(PlayerX)
	numChan <- 1

	FinishX(numChan, isDone)

}

func pinger(ping, pong, done chan *PlayerX, p *PlayerX, numchan chan int) {
	rand.NewSource(time.Now().UnixNano())

	for a := range ping {
		<-numchan
		a.Hit++
		a.Name = p.Name
		v := rand.Intn(100-1) + 1

		printPlayer(v, a)

		if v%BreakPoints == 0 {
			done <- a
			numchan <- v
			return
		}

		pong <- a
		numchan <- v
	}
}

func ponger(ping, pong, done chan *PlayerX, p *PlayerX, numchan chan int) {
	rand.NewSource(time.Now().UnixNano())

	for b := range pong {
		<-numchan
		b.Hit++
		b.Name = p.Name
		v := rand.Intn(100-1) + 1

		printPlayer(v, b)

		if v%BreakPoints == 0 {
			done <- b
			numchan <- v
			return
		}

		ping <- b
		numchan <- v
	}
}

func printPlayer(num int, p *PlayerX) {
	action := ThrowAction(p.Name)

	fmt.Println(p.Name, action, p.Hit, "counter ke", num)
	time.Sleep(time.Second)
}

func FinishX(numchan chan int, done chan *PlayerX) {
	for d := range done {
		action := ThrowAction(d.Name)
		for num := range numchan {
			fmt.Println(d.Name, "kalah pada", action, d.Hit, "counter ke", num)
			return
		}
	}
}

func ThrowAction(name string) string {
	var action string = ""
	switch {
	case name == "Player A":
		action = "ping->"
	default:
		action = "pong->"
	}
	return action
}
