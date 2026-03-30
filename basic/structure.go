package main
import "fmt"
func main(){
	type person struct{
		name string
		age int
		height float64
	
	}
	var info person = person{"Parth", 22, 5.9}
	fmt.Println(info)
}