package main
import "fmt"
func callByRefrence(str1,str2 *string) string{
	res:= *str1 + "___" +*str2
	return  res
}
func main(){
	var str1 = "who"
	var str2 = "are you"
	result := callByRefrence(&str1, &str2)
	fmt.Println(result)
	fmt.Println(&str1)
	fmt.Println(&str2)
}