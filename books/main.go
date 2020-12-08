package main

import (
	"fmt"
	"math/rand"
)

func add(a, b int) int {
	return a + b
}

const (
	_ = iota + 5
	a
	b
	c
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb
	gb
	tb
)

func oddEven(n int) {
	for i := 1; i < n; i++ {
		if i%2 == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}
}
func randomString() string {
	ans := ""
	for i := 0; i < 7; i++ {

		randomValue := (rand.Intn(100) % 26)

		ans = ans + string(rune(randomValue)+'A')
	}
	return ans
}
func main() {
	fmt.Println("form book main")

	// fmt.Println(add(1, 3))

	// x := 5
	// x++
	// fmt.Println(x)

	// fmt.Println(a, b, c)
	// fmt.Println(kb, mb, gb, tb)

	// slice := []int{}
	// for i := 1; i < 10; i++ {
	// 	slice = append(slice, i)
	// }
	// v := slice
	// for i := range slice {
	// 	fmt.Println(&slice[i], &v[i])
	// }

	// hell := make([][]string, 4)
	// for i := range hell {
	// 	hell[i] = make([]string, 4)
	// }

	// for i := 1; i < 4; i++ {
	// 	for j := 1; j < 4; j++ {
	// 		hell[i][j] = randomString()
	// 		fmt.Println(hell[i][j])
	// 	}
	// 	fmt.Println()
	// }

	// mp := make(map[string]int)
	// mp["hehhe"] = 1
	// mp["life of pi"] = 2

	// if _, ok := mp["abcd"]; !ok {
	// 	fmt.Println(ok)
	// }
	// fmt.Println(mp)

	el := make(map[string]map[string]int)
	fmt.Println(el)

	elements := map[string]map[string]int{
		"adress": map[string]int{
			"house": 1,
			"road":  2,
		},
		"dob": map[string]int{
			"date":  14,
			"year":  1996,
			"month": 4,
		},
	}
	fmt.Println(elements["adress"]["house"])
	fmt.Println(elements)

	x := make([]int, 3, 9)
	fmt.Println(x)

}
