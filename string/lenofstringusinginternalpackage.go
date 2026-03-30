package main
import "fmt"
import "strings"
func main(){
	var st string = "Iterate Over String Characters"
	fmt.Println(st)
	var result int =len(strings.TrimSpace(st))
	fmt.Println(result)
}