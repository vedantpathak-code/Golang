package main 
import(
	"fmt"
	"strings"
	"regexp" // helps to oprate pattern based string operations
)
func main(){
	str1:= "hello world welcome to golang"
    str2 := "golang is a programming language"
	fmt.Println(str1)
	fmt.Println(str2)
	res := regexp.MustCompile(`[em]`)
	finalres := res.Split(str1, -1)
	fmt.Println(finalres) 
	new := strings.Split(str1, " ")
	fmt.Println(new)
}	