package helper

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Data dari channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Data dari channel 2"
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2", data)
			counter++
		}
	if counter == 2{
		break
	}
	}
}

//deselect
func TestDeselectChannel(t *testing.T) {
	kopet1 := make(chan string)
	kopet2 := make(chan string)

	go func ()  {
		time.Sleep(1 * time.Second)
		kopet1 <- ("Data dari kopet1")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		kopet2 <- ("Data dari kopet2")
	}()

	counter := 0
	for{
		select{
		case data := <-kopet1:
			fmt.Println("Data Dari kopet1", data)
			counter++
		case data := <-kopet2:
			fmt.Println("Data dari kopet2", data)
			counter++
		default:
			fmt.Println("Menunggu kopet")
		}
		if counter == 2{
			break
		}
	}
}