package helper

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Loop ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println(data)
	}
	fmt.Println("Done")
}