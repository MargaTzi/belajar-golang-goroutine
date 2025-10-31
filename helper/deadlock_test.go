package helper

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	var user1 sync.Mutex
	var user2 sync.Mutex

	go func ()  {
		user1.Lock()
		fmt.Println("Gorutin 1: Lock")
		time.Sleep(1 * time.Second)

		user2.Lock()
		fmt.Println("Gorutin 1: Selesai")
		user2.Unlock()
		user1.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)

	go func() {
		user2.Lock()
		fmt.Println("Gorutin 2: Lock")
		time.Sleep(1 * time.Second)

		user1.Lock()
		fmt.Println("Gorutin 2: Selesai")
		user1.Unlock()
		user2.Unlock()
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}