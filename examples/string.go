package main
import "fmt"

func main(){
    s:="hello"
    fmt.Println(s)
    c:=[]rune(s)
    fmt.Println(c)
    c[0] = 'c'
    s2:=string(c)
    fmt.Println(s2)
}
