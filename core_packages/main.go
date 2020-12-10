// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	kids := []person{
// 		{"jill", 9},
// 		{"hehhe", 13},
// 		{"smith", 21},
// 		{"alas", 14},
// 	}
// 	sort.Sort(p(kids))
// 	fmt.Println(kids)
// }

// type p []person

// func (a p) Len() int           { return len(a) }
// func (a p) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a p) Less(i, j int) bool { return a[i].age < a[j].age }

// package main

// import "fmt"

// func main() {
// 	// h := crc32.NewIEEE()
// 	// h.Write([]byte("test"))
// 	// v := h.Sum32()
// 	// fmt.Println(v)
// 	a := " "
// 	fmt.Scanln(&a)

// 	fmt.Println(a)
// }

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	h.Write([]byte("kdkjfdkf 1"))
	bs := h.Sum([]byte{})
	fmt.Println(string(bs))
}
