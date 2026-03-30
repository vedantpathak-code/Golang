package main
import "fmt"
func Swap(a, b string) (string, string){
	return b,a
}
func main(){
	var a string = "fine!"
	var b string = "are you"
    swapa, swapb := Swap(a,b)
	fmt.Println("before swap:", a, b)
	fmt.Println("after swap:", swapa , swapb)
}