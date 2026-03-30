package main
import (
	"fmt"
	"strings"
)

func substring(str string, start , end int)string{
	return (strings.TrimSpace(str[start:end]))
}
func main(){
	var str string = "i am a string"
	fmt.Println(substring(str, 0, 10))
}