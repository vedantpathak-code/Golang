package main
import "fmt"
func main(){
	str1 := "hi im string"
	str2 := "i am a substring"
	if str1[:12]==str2{
		fmt.Println("str2 is equal to substring of str1")
	} else if str1[:10] <= str2{
		fmt.Println("str2 is less than or equal to substring of str1")
	} else if str1[:10] >= str2{
		fmt.Println("str2 is greater than or equal to substring of str1")
	}else{
		fmt.Println("str2 is not equal to substring of str1")
	}
	
}