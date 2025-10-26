package helper

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(10 * time.Second)
	channel <- "Golang Channel" //send
}

func TestChannelAsParameter(t *testing.T) {
	channnel := make(chan string)

	go GiveMeResponse(channnel)
	
	data := <-channnel //receive
	fmt.Println(data)
	close(channnel)
}

func OnlySend(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Wahyu Belajar Golang"
}

func OnlyReceive(channel <-chan string){
	data := <-channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string, 3)
	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	go OnlySend(channel)
	go OnlyReceive(channel)
	time.Sleep(3 * time.Second)
	close(channel)
}

func TestChannel(t *testing.T) {
	channel := make(chan string, 3)
	channel <- "Satu"
	channel <- "Dua"

	// cek kapasitas dan jumlah data saat ini
	fmt.Println("Capacity:", cap(channel)) // total slot = 3
	fmt.Println("Length:", len(channel))   // sudah terisi 2

	// kirim data ke-3
	channel <- "Tiga"

	fmt.Println("Length setelah tambah:", len(channel)) // sekarang penuh = 3

	// ambil satu data
	fmt.Println("Ambil data:", <-channel) // ambil "Satu"

	// sekarang 1 slot kosong lagi
	fmt.Println("Length sekarang:", len(channel))
}