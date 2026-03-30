//Go language does not provide the support of unions, but unions can be used as interface{} to hold any type of value
package main
import "fmt"
func main(){
	var uni interface{} = "Parth"
	fmt.Println(uni)
	uni = 22
	fmt.Println(uni)
	uni = "golang"
	fmt.Println(uni)
}