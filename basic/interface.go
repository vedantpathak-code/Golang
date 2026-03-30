package main
import "fmt"
type Animal interface{
	speak () string
}
type Dog struct{}
func (d Dog) speak() string{
	return "woof"
}
type Cat struct{}
func (c Cat) speak() string{
	return "meow"
}
func main(){
	var a Animal = Dog{
		
	}
	var b Animal = Cat{
		
	}
	fmt.Println(a.speak())
	fmt.Println(b.speak())
}