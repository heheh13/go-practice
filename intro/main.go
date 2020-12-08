package main

var (
	i int
	k float32 = 45
)

func main() {
	help()
	// fmt.Println("hello go!")

	// fmt.Println(i)
	// // k := 32.0
	// fmt.Println(k)

	// fmt.Printf("%v %T", k, k)
	// println()
	// var (
	// 	help string = "help"
	// 	name string = "name"
	// )
	// fmt.Println(help, name)

	// var j float32
	// j = 44.54
	// var num int = int(j)
	// fmt.Println(num)
	// i = 6500545
	// var j string
	// j = strconv.Itoa(i)
	// fmt.Println(j)

	// a := 10.0
	// b := 3.0

	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// // fmt.Println(a % b)

	// a := 10 //1010
	// b := 3  //0011

	// fmt.Println(a | b) //1011
	// fmt.Println(a & b) //0010
	// fmt.Println(a ^ b) //1001
	// fmt.Println(a &^ b)// 1000                   &^ and not  operator (a & ~b)

	// a := 8
	// fmt.Println(a << 1)
	// fmt.Println(a >> 1)

	// n := 3.14
	// n = 13.7e72
	// fmt.Println(n)

	// var n complex64 = 3i + 2i + 4i + 3.5
	// var m complex64 = 3 + 3i
	// fmt.Printf("%v %T \n", n, n)
	// fmt.Println(n + m)

	//real(vlaue), complex(r,im)

	// s := "this is a string"
	// // s1 := "heehhe"
	// // fmt.Printf("%v %T \n", string(s[2]), string(s[2]))
	// // fmt.Println(s + s1)

	// b := []byte(s) //convert stirngs into unit8
	// fmt.Printf("%v , %T \n", b, b)

	// r := 'a'
	// fmt.Printf("%v %T", r, r)

	// constant

	//arrays

	// grades := [...]int{1, 2, 3}
	// fmt.Printf("grades : %v ", grades)

	// student := [...]string{"mehedi hasan", "hehhe", "asdf"}
	// fmt.Println(student)

	//multidimentional arrays

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Printf("%v %T\n", a, a)
	// fmt.Printf("len : %v cap : %v \n", len(a), cap(a))

	// b := a
	// b[1] = 0 ///changes reflects as pointers
	// fmt.Printf("%v \n, %v \n", a, b)

	// a := make([]int, 3, 100)
	// fmt.Println(a)
	// a = append(a, 1)
	// fmt.Println(a)

	// statePopulation := map[string]int{
	// 	"ban": 123,
	// 	"usa": 5454,
	// 	"ind": 545,
	// }

	// if pop, ok := statePopulation["ban"]; ok {
	// 	fmt.Println(pop)
	// }

	// number := 50
	// guess := 70

	// if guess < number {
	// 	fmt.Println("too low")
	// } else if guess > number {
	// 	fmt.Println("too high")
	// } else {
	// 	fmt.Println("wkkwkwkwkwkkwkwk")
	// }

	// fmt.Println(number <= guess, number >= guess, number != guess)

	// for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	// for i := 1 , i <= 10; i = i+1 {
	// 	for j := 1 , j <= 10; j++ {
	// 		fmt.Println(i*j)
	// 	}
	// }

	// Loop:

	// 	for i := 1; i < 10; i++ {
	// 		for j := 1; j < 10; j++ {
	// 			println(i * j)
	// 			if i*j > 50 {

	// 				break Loop
	// 			}
	// 		}
	// 	}

	// s := []int{1, 2, 3, 4}
	// for k := range s {
	// 	fmt.Println(s[k])

	// }

	//pointers

	// a := 44
	// var b *int = &a
	// fmt.Println(a, *b)
	// a = 27
	// fmt.Println(a, *b)
	// *b = 45
	// fmt.Println(a, *b)

	// a := []int{3, 3, 1, 2, 3, 5}
	// for i := range a {
	// 	fmt.Println(&a[i])
	// }

	// a := []int{1, 2, 3}
	// for i := 0; i < 3; i++ {
	// 	a = append(a, i)
	// }
	// fmt.Println(a)

	// fun := func() {
	// 	fmt.Println("hello anynomous func")
	// }
	// fun()

	// interfaces

	// var w Writer = ConsoleWriter{}
}

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct{}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }
