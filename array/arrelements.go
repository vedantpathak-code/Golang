package main
import "fmt"
func main(){
	arr := [8] string{"apple", "banana", "orange", "mango", "grape", "pineapple"}
	fmt.Println("element of array: ")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}