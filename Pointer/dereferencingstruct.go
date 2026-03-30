package main 
import "fmt"
type Student struct {
	name string
	age int 

}
func main(){
	s1 := &Student{"Parth", 22}
	fmt.Printf("student name: %s\n student age: %d", s1.name, s1.age)


}