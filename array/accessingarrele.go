package main 
import "fmt"
func main(){
	var n [5]int
	var i , j int
	for i = 0; i < 5; i++{
		n[i] = i + 1
		fmt.Println("for i :",n[i]) 
	} 
	for j = 0; j < 5; j++{
		n[j] = n[j] * 2
		fmt.Println("for j :",n[j])
	}
}