package main
import "fmt"
func main(){
	var q int = 92
	var ptr *int = &q // &q address of q
	fmt.Println(*ptr)
}