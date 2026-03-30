package main 
import "fmt"
func main(){
	var arr [3][4] string = [3][4] string {
		{"apple", "banana", "orange", "mango"},
		{"grape", "pineapple", "kiwi", "watermelon"},
		{"peach", "pear", "plum", "strawberry"},
	}
	fmt.Println(arr)
	fmt.Println(arr[2][3])
}	