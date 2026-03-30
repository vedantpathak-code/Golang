package main
import "fmt"
func main(){
	n := []int{1,2,3,4,5,6,7,8,9}
	m := 5
	for _, num := range n {
		if num == m {
			fmt.Println("found :", num)
			break
		} 
		fmt.Println("current number", num)

	}
	fmt.Println("loop ended")
} 
