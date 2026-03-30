package main
import "fmt"
func main(){
	x := 10
	ptr1 := &x
	y := 19
	ptr2 := &y
	fmt.Printf("value of x: %d, value of y: %d\n", *ptr1, *ptr2)
	fmt.Printf("address of x: %s, address of y: %s\n", ptr1, ptr2)
}