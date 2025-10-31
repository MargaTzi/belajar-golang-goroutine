package repository

import "fmt"

func Deadlock() {
	channel := make(chan string)

	channel <- "Halo"

	fmt.Println(<-channel)
}