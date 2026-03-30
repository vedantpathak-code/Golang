package main
import "fmt"
import "strings"
func main(){
	str1 := "hello"
	str2 := "HELLO!"
	if strings.EqualFold(str1, str2){
		fmt.Println("Strings are equal")
	}else{
		fmt.Println("Strings are not equal")
	}
}