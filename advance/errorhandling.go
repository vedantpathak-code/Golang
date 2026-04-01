/*package main
import(
	"errors"
	"fmt"
	"math"
)
func Sqrt(value float64) (float64, error){
	if (value <0){
		return 0, errors.New("Math: Negative number")
	}
	return math.Sqrt(value),nil
} 
func main(){
	result,err :=Sqrt(16)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	} 
	
}*/
package main

import (
	"fmt"
	"time"
)

func display(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	go display("Hello, Goroutine!") 
	display("Hello, Main!")         
	time.Sleep(time.Second) 
}

