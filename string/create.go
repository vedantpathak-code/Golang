package main 
import "fmt"
import "reflect"
func main(){
	var string1 string = "hello"
	fmt.Println(string1)
	fmt.Println(len(string1))
	fmt.Println(string1[0:3])
	fmt.Println(reflect.TypeOf(string1))
}