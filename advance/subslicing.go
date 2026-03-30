package main
import "fmt"
func slice(x []int){
	fmt.Printf("len=%d cap=%d slice%v\n",len(x),cap(x),x)
}
func main(){
	numbers :=[10]int{23,32,543,42,242,42,6,54,84,90}
	subslice1 := numbers[3:5]
	subslice2 := numbers[5:9]
	slice(subslice1)
	slice(subslice2)
}