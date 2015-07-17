package main
import "fmt"

func main(){

    c := 'f'

    switch {
    case '0' <= c && c <= '9':
        fmt.Println(c - '0')
    case 'a' <= c && c <= 'f':
        fmt.Println(c - 'a' + 10)
    case 'A' <= c && c <= 'F':
        fmt.Println(c - 'A' + 10)
    default:
        fmt.Println(0)
    }
}
