package main

import (
	"fmt"
	"sync"
)

// func help() {

// 	fmt.Println("from paralel")

// 	go sayHello()

// 	time.Sleep(100 * time.Millisecond)
// }

// func sayHello() {
// 	fmt.Println("Hello")

// 	fmt.Printf("%v \n", runtime.GOMAXPROCS(-1))
// }

var wg = sync.WaitGroup{}

func help() {
	x := 30
	ch := make(chan int, x)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()

}
