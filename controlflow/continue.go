package main
import "fmt"
func main(){
	for i :=0 ; i < 11 ; i++ {
		if i == 7 {
			continue
		}
	fmt.Println(i)	
}
}