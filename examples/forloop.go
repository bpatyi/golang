package main
import "fmt"

func main(){
    for i:=0; i < 10; i++{
        if i > 5{
            break
        }
        fmt.Println(i)
    }

    fmt.Println("=========================")

    J: for j := 0; j < 5; j++ {
        for i := 0; i < 10; i++{
            if i > 5 {
                break J
            }
            fmt.Println(i)
        }
    }

    fmt.Println("=========================")

    for i := 0; i < 10; i++{
        fmt.Println("%d", i)
    }

    fmt.Println("=========================")

    i := 0

    Loop:
        if i < 10{
            fmt.Println("%d", i)
            i++
            goto Loop
        }
}
