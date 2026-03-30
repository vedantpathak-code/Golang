package main
import "fmt"
func main(){ 
	fmt.Println("starting")
    defer fmt.Println("defer function will executed")
    fmt.Println("ending")	
}