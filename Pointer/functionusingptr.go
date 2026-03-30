package main 
import "fmt"
func FunctionUsingPtr(num *int) {
	*num = *num * 30
}
func main(){
	number := 30
	FunctionUsingPtr(&number)
	fmt.Println(number)
}