package main
import "fmt"
import "strings"
func main(){
	string1 := "hello11"
	string2 := "hello111111"
	result := strings.Compare(string1, string2)
	fmt.Println(result)
}