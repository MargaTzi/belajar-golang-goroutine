package helper

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(1 * time.Second)

	tick := <-channel
	fmt.Println(tick)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("execute after 1 second")
		group.Done()
	})

	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}
}