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

## rand

        math/rand
        rand.Seed(n)
        rand.Intn(n)

## notes

    go is staticly typed
    very serious oabout typeconversion
    
    str += string(rune(n)+ 'a')
