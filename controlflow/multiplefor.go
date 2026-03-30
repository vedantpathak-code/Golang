package main
import "fmt"
func main(){
	var a int = 20
	var b int
	numbers := [5]int{1,2,3,4,5,}
	for b := 0; b < 10; b++{
		fmt.Println("b:",b)
	}
	for b < a{
		b++
		fmt.Println("b:",b)
	}
	for i,x:= range numbers{
		fmt.Println("number:",numbers[i],x)
	}

	}

