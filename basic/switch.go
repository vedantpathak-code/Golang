package main
import "fmt"
func main(){
	day := "Monday"
	switch day{
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	case "Wednesday":
		fmt.Println("It's Wednesday!")
	default:
		fmt.Println("It's another day of the week.")			
	}
}	