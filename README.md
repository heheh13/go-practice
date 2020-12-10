# go-practice

## go run

need to build all files if there are multiples go files in same directory

cd dir -> go build .
./dirname

go run -race src/main.go

## constant

`iota` is a special case that is used for enumrate constant

    * imutable
    * must be callcuated at complied time
    * upper case constant exported to package level
    * can use bitwise op to initiate const

## looping

    * for i ; i < n ;i ++{}
    * for {}
    * for i , j = 0 , 0 ; cond ; i , j = i+1 , j+1{}
    * can use break
    * can use labels to exit from loop

## arrays

        can't declared variable size
        can be copy to a new one with values

## slice

        work with ref
        slice = append(slice, val)
        
        having fun with mulitdimentional slice
        
        hell := make([][]int, 4)
        
        for i := range hell {
            hell[i] = make([]int, 4)
        }

## map

        if _, ok := mp["abcd"]; !ok {
            fmt.Println(ok)
        }

        m := make(map[string]int)
        m := map[string]int{
            "number":1,
        }

        multidimentainal maps

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

## functions

        func heheh(name string, age int) (string, int) {
            return name, age
        }

        func vaiadic(values ...int) {
            fmt.Println(values)
        }

`clousers`

        func everGenerator() func() int {
            x := 0
            return func() (ret int) {
                ret = x
                x += 2
                return ret
            }
        }

        nextEven := everGenerator()
        fmt.Println(nextEven())
        fmt.Println(nextEven())
        fmt.Println(nextEven())

`recursion`

        func fact(n int) int {
            if n == 0 {
                return 1
            }

            return fact(n-1) * n
        }
`defer`

    call after function returns?

    linfastout

    can use to some thing that need to close letter

    Has a usecase in recovering panic

`panic and recover`

        func panicAndRecover() {
                defer func() {
                    str := recover()
                    fmt.Println(str)
                    // panic("im still panicing")
                }()
                panic("im panicing")
            }
        }
        need some more understanding about recover

## pointers

        a := 42
        var b *int = &a
        fmt.Println(a, b)
        fmt.Println(a, *b)
        fmt.Printf("%d %d ", &a, &b)

## structs

`decleartions of structs`

        type Doctor struct {
            number     int
            actorName  string
            companions []string
        }

        aDoctor := Doctor{
                number:    3,
                actorName: "jhon partwee",
                companions: []string{
                    "heheh hehe ",
                    "jo grant",
                    "shara jane",
                },
            }
        fmt.Println(aDoctor)

`struct can be embeded to over come inheritence`

        type animal struct {
            name   string
            origin string
        }

        type bird struct {
            animal
            speed  int32
            canFly bool
        }

        b := bird{
            animal: animal{
                name:   "doyel",
                origin: "bd",
            },
            canFly: true,
            speed:  60,
        }

        fmt.Println(b)

`slices of structures`

        type point struct {
            x, y int
        }

        points := []point{}

        for i := 0; i < 10; i++ {
            p := point{x: i, y: i}
            points = append(points, p)
        }
        fmt.Println(points)
`structure methods`

        type Circle struct {
            x, y, r float32
        }

        func (c *Circle) area() float32{
            return (math.Pi * c.r * c.r)
        }

        c := new(Circle)
        *c = Circle{0, 0, 5}
        fmt.Println(c.area())

## interfaces

`interfaces can be used as types where ever i want to use as a params type , return type , in  map ??`

`any shape having the area methods implemts the interfaces`

        type shape interface {
            area() float32
        }

        type shape2 interface {
            haveFun() float32
        }

        func (c *circle) haveFun() float32 {
            fmt.Println("having fun wiht circle")
            return math.Pi
        }
        func (r rect) area() float32 {
            return r.x * r.y
        }
`from main`

        c := circle{0, 0, 5}
        r := rect{2, 3}

        shapes := []shape{&c, r}

        for _, shape := range shapes {
            fmt.Println(shape.area())
        }
`struct circle and rect implements a different interfaces`

        difShape := []shape2{c, r}
        for _, shape := range difShape {
            fmt.Println(shape.haveFun())
        }

## concurrency

`will modify letter a lot to know`

- `recive only channel ch chan <- string`

- `send only channel ch <-chan string`

- `channels waits for reciving a values from a sender channnel`

`what this program do?`

`the fuction count and bellow for loop execute concurently`

`count function sends data to the cannel and loop section recive data form the channel`

    func count(str string, ch chan<- string) {
        for i := 0; i < 5; i++ {
            ch <- str
            time.Sleep(500 * time.Millisecond)
        }
        close(ch)
    }

    func main() {
        ch := make(chan string)
        go count("sheep", ch)
        for msg := range ch {
            fmt.Println(msg)
        }

    }

`select`

`allows us to first come first serve`

        func main() {
            ch1 := make(chan string)
            ch2 := make(chan string)

            go func() {
                for {
                    ch1 <- "every 500 ms"
                    time.Sleep(500 * time.Millisecond)
                }
            }()
            go func() {
                for {
                    ch2 <- "every 2000 ms"
                    time.Sleep(2000 * time.Millisecond)
                }
            }()
            for {
                select {
                case msg1 := <-ch1:
                    fmt.Println(msg1)
                case msg2 := <-ch2:
                    fmt.Println(msg2)
                }
            }
`concurrent?`

`what we have tried to do is calculate fibnacci using multiple core?`
``

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

## Packeges

`Packages are easy and simple to use`
`what we need to customise our own realative import and do our best usages`
`we can add aliases to avoid conflicts`

## Testing

`unit testing`

`write test for every functoins`

`usualy in someting_test.go`
`follow the convention`

## rand

        math/rand
        rand.Seed(n)
        rand.Intn(n)

## notes

`go is staticly typed`

`very serious oabout typeconversion`

`str += string(rune(n)+ 'a')`

`go run -race to check for data race`
