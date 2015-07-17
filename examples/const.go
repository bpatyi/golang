package main

import "fmt"

func main() {
    const(
        a = iota
        b
    )

    fmt.Println(a)
    fmt.Println(b)
}
