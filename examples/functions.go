package main
import "fmt"

func rec(i int){
    if i == 10 {
        return
    }
    rec(i+1)
    fmt.Println("%d", i)
}

var a int

func main(){

    rec(1)

    fmt.Println("================")

    a = 5
    print(a)
    f()
}

func f(){
    a := 6
    print(a)
    g()
}

func g(){
    print(a)
}
