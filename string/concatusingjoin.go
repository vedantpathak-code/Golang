package main
import "fmt"
import "strings"
func main(){
	str1 := []string{"hello", "world", "golang"}
	str2 := []string{"welcome", "to", "golang"}
	res1 := strings.Join(str1, "-") 
	res2 := strings.Join(str2,"*")
	fmt.Println(res1)
	fmt.Println(res2)
	
}