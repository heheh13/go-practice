package main

import (
	"fmt"

	"github.com/heheh13/go-practice/packages/utility"
	heheh "github.com/heheh13/go-practice/packages/utility/helper"
)

func main() {
	fmt.Println("packages")
	fmt.Println(utility.Hello())
	fmt.Println(heheh.Hello())
	fmt.Println(utility.Adder(1, 2, 3, 4))
}
