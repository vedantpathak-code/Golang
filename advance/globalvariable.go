package main
import "fmt"
var global_variable int
func main (){
	a := 33
	b := 74
	global_variable = a + b
	fmt.Println("addition using global variable", global_variable)
}