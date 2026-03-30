package main
import "fmt"
import "reflect"
func main(){
	var arr [10]int = [10]int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr))
}