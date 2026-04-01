package main
import "fmt"
func factorial (i int) int{
	if (i<=1){
		return i
	}
	return i *factorial (i-1)
}

func main(){
	i := 18
	fmt.Println(factorial(i))


}