package main
import "fmt"
func main(){
	dayOfWeek := 6
	switch dayOfWeek{
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thrusday")
	case 5: 
		fmt.Println("friday")
	default:
		fmt.Println("invalid!")	
	}
 }