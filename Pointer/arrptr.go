package main
import "fmt"
func main(){
	i := [3] int{10,20,30}
	var a int
	var ptr [3]*int
	for a = 0; a < len(i); a++{
		ptr[a] = &i[a]
	}
    for a = 0; a <len(i) ; a++{
		fmt.Printf("index %d: %d\n", a, *ptr[a])
	}
}