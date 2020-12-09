package main

import "fmt"

//recive only channel ch chan <- string
// send only channel ch <-chan string

// func count(str string, ch chan<- string) {
// 	for i := 0; i < 5; i++ {
// 		ch <- str
// 		// time.Sleep(500 * time.Millisecond)
// 	}
// 	close(ch)
// }
// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		for {
// 			ch1 <- "every 500 ms"
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}()
// 	go func() {
// 		for {
// 			ch2 <- "every 2000 ms"
// 			time.Sleep(2000 * time.Millisecond)
// 		}
// 	}()
// 	for {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println(msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}
