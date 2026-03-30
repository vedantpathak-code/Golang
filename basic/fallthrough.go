package main
import "fmt"
func main(){
	day := "Saturday"
	switch day{
	case "Monday":
		fmt.Println("it is Monday")
	case "Tuesday":
		fmt.Println("it is Tuesday")
	case "Wednesday":
		fmt.Println("it is Wednesday")
	case "Thursday":
		fmt.Println("it is Thursday")
	case "Friday":
		fmt.Println("it is Friday")
	case "Saturday":
		fmt.Println("it is Saturday")
		fallthrough
	case "Sunday":
		fmt.Println("it is Sunday")
	default:
		fmt.Println("invalid day")

	}
}