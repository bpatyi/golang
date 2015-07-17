package main
import "fmt"

func Panic(f func()) (b bool)  {
    defer func() {
        if x := recover(); x != nil {
            b = true
        }
    }()
    f()
    return
}

func panicy()  {
    // panic == true
    // var a []int
    // a[3] = 5

    // panic == false
    var a [3]int
    a[2] = 5
}

func main()  {
    fmt.Println(Panic(panicy))
}
