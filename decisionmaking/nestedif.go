package main
import "fmt"
func main(){
	a := true
	b := false
	if a == true{
		fmt.Println("first loop executed")
		if b == false{
			fmt.Println("second loop executed")
		}
	}
}