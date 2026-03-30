package main
import "fmt"
func main(){
	var new bool = true
	var old bool = false
	fmt.Println("new:", new)
	fmt.Println("old:", old)

	if new{
		fmt.Println("it is a true")
	} else{
		fmt.Println("it is a false")
	}
}