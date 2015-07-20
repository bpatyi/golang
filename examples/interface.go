package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}
type Cat struct{}
type Llama struct{}
type JavaProgrammer struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

func (c Cat) Speak() string  {
    return "Meow!"
}

func (l Llama) Speak() string {
    return "?????"
}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

func main()  {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
