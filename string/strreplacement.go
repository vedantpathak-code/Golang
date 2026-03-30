package main
import "fmt"
import "strings"
func main(){
	str1 := "replace"
	str2 := "need to replace"
	res1:= strings.Replace(str1, "r", "R", 1)
	res2 := strings.Replace(str2, "e", "E" ,-1)
	fmt.Println(res1)
	fmt.Println(res2)
}