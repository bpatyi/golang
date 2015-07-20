package main
import "fmt"

func average(xs *[4]float64) (avg float64) {
    sum := 0.0
    switch len(xs) {
    case 0:
        avg = 0
    default:
        for _, v := range xs {
            sum += v
        }
        avg = sum / float64(len(xs))
    }
    return
}

func main(){
    arr := [4]float64{1.1, 2.0, 3.0, 5.4}

    avg := average(&arr)
    fmt.Println(avg)
}
