package main

import "fmt"

func square() func() int {
count := 0
return func() int {
count++
return count * count
}
}

func main() {
squareValue := square()
fmt.Println(squareValue())
fmt.Println(squareValue())
fmt.Println(squareValue())
}
