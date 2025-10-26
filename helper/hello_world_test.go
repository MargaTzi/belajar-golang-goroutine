package helper

import (
	"fmt"
	"testing"
	"time"

)


func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T){
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T){
	for i := 0; i < 100000; i++{
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}

func DisplayMessage(message string){
	fmt.Println("Pesan", message)
}

func TestManyMessage(t *testing.T){
	for i:= 0; i < 10000; i++{
		go DisplayMessage("Belajar Golang Routine ke-" + fmt.Sprint(i))
	}
	time.Sleep(2 * time.Second)
}