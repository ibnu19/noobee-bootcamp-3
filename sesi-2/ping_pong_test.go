package sesi2concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Player struct {
	Name string
	Hit  int
}

const BreakPoints int = 11

func TestPingPongApp(t *testing.T) {
	ping := make(chan *Player)
	pong := make(chan *Player)
	isDone := make(chan *Player)
	numChan := make(chan int)

	A := Player{Name: "Player A", Hit: 0}
	B := Player{Name: "Player B", Hit: 0}

	// Player A
	go pinger(ping, pong, isDone, &A, numChan)
	// Player B
	go ponger(ping, pong, isDone, &B, numChan)

	ping <- new(Player)
	numChan <- 1

	Finish(numChan, isDone)

}

func pinger(ping, pong, done chan *Player, p *Player, numchan chan int) {
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

func ponger(ping, pong, done chan *Player, p *Player, numchan chan int) {
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

func printPlayer(num int, p *Player) {
	action := ThrowAction(p.Name)

	fmt.Println(p.Name, action, p.Hit, "counter ke", num)
	time.Sleep(time.Second)
}

func Finish(numchan chan int, done chan *Player) {
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
