package main
import "fmt"
func main(){
	var q int = 92
	var ptr *int = &q
	fmt.Println(*ptr)
	fmt.Println(&q)
}