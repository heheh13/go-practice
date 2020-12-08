package main

import "fmt"

func heheh(name string, age int) (string, int) {
	return name, age
}

func vaiadic(values ...int) {
	fmt.Println(values)
}

func everGenerator() func() int {
	x := 0
	return func() (ret int) {
		ret = x
		x += 2
		return ret
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return fact(n-1) * n
}

func testDefer() {
	nextEven := everGenerator()
	fmt.Println(nextEven())
	defer fmt.Println(nextEven())
	fmt.Println(nextEven())
}
func panicAndRecover() {
	defer func() {
		str := recover()
		fmt.Println(str)
		panic("im still panicing")
	}()

	panic("im panicing")
}

func main() {
	fmt.Println("functions")
	fmt.Println(heheh("mehedi", 24))

	vaiadic(1, 2, 3, 4, 5)

	//clouser & annynoums function

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(10, 15))

	nextEven := everGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(fact(5))

	testDefer()

	panicAndRecover()
}
