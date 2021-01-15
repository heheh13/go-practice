package main

import "fmt"

//Dog is structure
// will be used for store a  collection of dog
type Dog struct {
	age  int
	name string
}

func main() {
	// alpha := Dog{
	// 	age:  1,
	// 	name: "alpha",
	// }
	// beta := Dog{
	// 	age:  2,
	// 	name: "beta",
	// }
	// fmt.Printf("%p %p\n", &alpha, &beta)
	// dogs := []Dog{alpha, beta}
	// fmt.Printf("%p %p\n", &dogs[0], &dogs[1])
	// for i := 0; i < 2; i++ {
	// 	fmt.Printf("%p\n", &dogs[i])
	// }
	// for _, i := range dogs {
	// 	fmt.Printf("%p\n", &i)
	// }

	// arr := []int{}
	// n := 3
	// for i := 0; i < n; i++ {
	// 	arr = append(arr, i)
	// 	fmt.Printf("%d %d\n", arr[i], &arr[i])
	// }
	// arr[1] = 32
	// for i := range arr {
	// 	fmt.Println(i)
	// }
	// fmt.Println(arr)

	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// time.Sleep(1 * time.Second)

	// s := []int{1, 2, 3}
	// b := make([]int, len(s))
	// copy(b, s)
	// b[1] = 33
	// fmt.Println(s, b)
	// s = append(s[:1], s[:1]...)
	// fmt.Println(s)
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// i, j := 3, 7
	// fmt.Println(a[i:], a[j:])
	// copy(a[i:], a[j:])
	// // for k, n := len(a)-j+i, len(a); k < n; k++ {
	// // 	fmt.Println("here")
	// // 	a[k] = 0 // or the zero value of T
	// // }
	// a = a[:len(a)-j+i]
	// fmt.Println(a, len(a), cap(a))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slidingWindow(5, a))

}

func slidingWindow(size int, input []int) [][]int {
	// returns the input slice as the first element
	if len(input) <= size {
		return [][]int{input}
	}

	// allocate slice at the precise size we need
	r := make([][]int, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}
