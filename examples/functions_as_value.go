package main
import "fmt"

func main(){
    a := func() {
        fmt.Println("Hello")
    }

    a()

    var xs = map[int]func() int{
        1: func() int { return 10 },
        2: func() int { return 20 },
        3: func() int { return 30 },
    }

    fmt.Println(xs)
}
