package main 
import "fmt"
func main(){
	var calculate = func( a , b , c int) int{
		return a + b+  c
	}
	fmt.Println(calculate(10,4,9))
}