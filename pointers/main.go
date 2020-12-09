package main

import "fmt"

func main() {
	fmt.Println("pointers")
	a := 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)
	fmt.Printf("%d %d \n", &a, &b)
	a = 32
	fmt.Println(a, *b)

	arr := []int{1, 3, 4, 5}

	c := &arr

	x := arr
	fmt.Println(c)
	fmt.Println(x)
}
