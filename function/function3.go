package main 
import "fmt"

func main(){
	var a int= 200
	var b int= 300
	var res int
	res = Avrage(a, b)
	fmt.Println( res )
}
func Avrage(a,b int) int{
	return (a+b)/2
}
