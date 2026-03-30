package main
import "fmt"
func main(){
	q := 48
	ptr := &q
	fmt.Println("Value of q: ", *ptr)
	fmt.Println("Address of q: ", &ptr)
	*ptr = 100
	fmt.Println("Updated value of q: ", q)
}