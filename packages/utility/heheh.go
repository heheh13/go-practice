package utility

// Hello hi
// a custom func to learn how to create package
func Hello() string {
	return "hello from utility"
}
// Adder hi
// a custom func to learn how to create package
func Adder(args ...int) (sum int) {
	for i := range args {
		sum += i
	}
	return
}
