package main
import "fmt"
func main(){
	str1 := "Iterate Over String Characters"
	for index , char := range str1{
		fmt.Println( index, string(char))
		
	}

}