package helper

import (
	"fmt"
	"sync"
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

func TestRaceCondition(t *testing.T) {
	var x = 0
	for i := 0; i <= 1000; i++{
		go func() {
			for j := 1; j <= 100; j++{
				x = x + 1
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Count : ", x)
}

func TestMutex(t *testing.T){
	var x = 0
	var mutex sync.Mutex
	for i := 0; i <= 1000; i++{
		go func() {
			for j := 1; j <= 100; j++{
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Count : ", x)
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i :=0; i<100; i++{
		go func() {
			for j :=1; j< 100; j++{
				account.addBalance(1)
				fmt.Println(account.getBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.getBalance())
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name: "Dito",
	}
	user2 := UserBalance{
		Name: "Wahyu",
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(5 * time.Second)
}

func TestWaitGroup(t *testing.T){
	group := &sync.WaitGroup{}

	for i := 0; i < 5; i++{
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}

func TestPool(t *testing.T){
	var pool = sync.Pool{
		New: func() interface{} {
			return "Default"
		},
	}
		
	
	pool.Put("Cario")
	pool.Put("Mada")
	pool.Put("Ardisto")

	for i := 0; i<10; i++{
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
}

func TestMap(t *testing.T){
	var data sync.Map
	var Addmap = func (value int)  {
		data.Store(value, value)
	}

	for i := 0; i < 100; i++{
		go Addmap(i)
	}

	time.Sleep(3 * time.Second)

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}