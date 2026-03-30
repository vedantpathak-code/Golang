package main 
import "fmt"
func main(){
	var x int
	var ptr *int
	x = 10
	ptr = &x
	*ptr = 20
	fmt.Println(x)
	value := *ptr
	fmt.Println(value)
	
}