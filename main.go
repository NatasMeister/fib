package main

import (
    "fmt"
    "github.com/NatasMeister/fib/fib"
)

func main() {
    fmt.Println("Fantastiske Fibonacci!")

    n := 10
    result := fib.Fib(n)
    fmt.Println("Fib(", n, ") = ", result)
}