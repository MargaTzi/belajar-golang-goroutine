package helper

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDeadlockPasti(t *testing.T) {
	var mu1 sync.Mutex
	var mu2 sync.Mutex

	start := make(chan struct{}) // buat nyinkronin dua goroutine

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1: lock mu1")

		// kasih sinyal bahwa goroutine 1 sudah lock mu1
		start <- struct{}{}

		// tunggu goroutine 2 juga sudah lock mu2
		<-start

		mu2.Lock()
		fmt.Println("Goroutine 1: lock mu2") // tidak akan pernah tercetak

		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu2.Lock()
		fmt.Println("Goroutine 2: lock mu2")

		// tunggu sinyal dari goroutine 1
		<-start

		// kirim sinyal balik agar goroutine 1 lanjut
		start <- struct{}{}

		mu1.Lock()
		fmt.Println("Goroutine 2: lock mu1") // tidak akan pernah tercetak

		mu1.Unlock()
		mu2.Unlock()
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("=== SELESAI ===")
}