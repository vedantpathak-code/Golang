package main 
import "fmt"
func SwapingByptr(a *int, b *int) {
	temp := *a
*a = *b
*b = temp

	
}



func main(){
	var num1 int = 29
	var num2 int = 32
	SwapingByptr(&num1, &num2)
	fmt.Println("After Swaping: ", num1, num2)
}
