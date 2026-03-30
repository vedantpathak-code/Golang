package main
import "fmt"
func main(){
	year := 2009
	if year%4==0 && year%100!=0 {
		fmt.Println(year, "it's a leap year")
	} else {
		fmt.Println(year, "it's not a leap year")
	}
}