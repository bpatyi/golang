package main

import (
    "even"
    "fmt"
)

func main()  {
    i := 5
    fmt.Println("Is %d even? %v", i, even.Even(i))

    // dead, because odd is a private function of even
    // fmt.Println("Is %d even? %v", i, even.odd(i))
}
