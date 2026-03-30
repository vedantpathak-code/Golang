package main
import "fmt"
func slice(x []int){
	fmt.Printf("len%d cap%d value%v\n",len(x),cap(x),x)
}
func main(){
	var number []int
	slice(number)
	number = append(number,3)
	slice(number)
	number = append(number,7)
	slice(number)
	number = append(number,4,6,9)
	slice(number)
	numbers1 := make([]int, len(number), cap(number)*2)
	copy(numbers1, number)
	slice(numbers1)

}