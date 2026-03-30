package main
import "fmt"
import "time"
func sayHello(){
	fmt.Println("Hello")

}
func main(){
	go sayHello()
	time.Sleep(time.Second)
}