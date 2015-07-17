package main
import "fmt"

func printit(x int){
    fmt.Println(x)
}

func callback(y int, f func(int)){
    f(y)
}

func main(){
    callback(10, printit)
}
