package main
import "fmt"

func main(){
    a := [...]int{1,2,3,4,5}
    s1 := a[2:4]
    fmt.Println(s1)
    s2:=a[1:5]
    fmt.Println(s2)
    s3:=a[:]
    fmt.Println(s3)
    s4:=a[:4]
    fmt.Println(s4)
    s6:=a[2:4:5]
    fmt.Println(s6)

    fmt.Println()
    fmt.Println("============")
    fmt.Println()

    s0 := []int{0,0}
    fmt.Println(s0)
    s7 := append(s0, 2)
    fmt.Println(s1)
    s8 := append(s7, 3, 5, 7)
    fmt.Println(s2)
    s9 := append(s8)
    fmt.Println(s9)

    fmt.Println()
    fmt.Println("============")
    fmt.Println()

    var b = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    var s = make([]int, 6)

    fmt.Println(s)

    n1:=copy(s, b[0:])
    fmt.Println(n1)
    fmt.Println(s)
    n2:=copy(s, s[2:])
    fmt.Println(n2)
    fmt.Println(s)

    fmt.Println()
    fmt.Println("============")
    fmt.Println()

    monthdays := map[string]int{
        "Jan": 31, "Feb": 28, "Mar": 31,
        "Apr": 30, "May": 31, "Jun": 30,
        "Jul": 31, "Aug": 31, "Sep": 30,
        "Oct": 31, "Nov": 30, "Dec": 31,
    }

    fmt.Println(monthdays)

    year := 0
    for _, days := range monthdays {
        year += days
    }
    fmt.Printf("Numbers of days in a year: %d\n", year)
}
