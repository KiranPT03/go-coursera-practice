package main

import (
	"fmt"
	"sync"
	"time"
)

func wait10sec(wg *sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	fmt.Println("Finished task of 10 sec")
	wg.Done()
}

func wait6sec(wg *sync.WaitGroup) {
	time.Sleep(6 * time.Second)
	fmt.Println("Finished task of 6 sec")
	wg.Done()
}

func wait3sec(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Finished task of 3 sec")
	wg.Done()
}

func main() {
	fmt.Println("Starting main goroutine")
	var wg sync.WaitGroup
	wg.Add(3)

	go wait10sec(&wg)
	go wait6sec(&wg)
	go wait3sec(&wg)

	wg.Wait()
	fmt.Println("Finished main goroutine")
}
