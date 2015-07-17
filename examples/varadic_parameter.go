package main
import "fmt"

func main()  {
    arg := map[string]int{
        "Jan": 31, "Feb": 28, "Mar": 31,
        "Apr": 30, "May": 31, "Jun": 30,
        "Jul": 31, "Aug": 31, "Sep": 30,
        "Oct": 31, "Nov": 30, "Dec": 31,
    }

    for _, n := range arg {
        fmt.Println("And the number is: %d\n", n)
    }
}
