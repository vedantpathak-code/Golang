package main
import "fmt"
type employee struct{
	name string
	age int
	salary int
	id string
}
func callByReference(emp *employee){
	emp.name = "Parth"
	emp.age = 22
	emp.salary = 50000
	emp.id = "EMP001"	
}
func main(){
	emp := employee{
		name: "John",
		age: 30,
		salary: 60000,
		id: "EMP002",
	
	} 
	fmt.Println("Before calling function:",emp)
	callByReference(&emp)

	fmt.Println("After calling function:",emp)
}