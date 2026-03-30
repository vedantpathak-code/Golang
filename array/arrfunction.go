package main 
import "fmt"
func arrAvg(arr[10] int) int{
	var i int 
	var sum int 
	var avg float64
	for i = 0; i < 10; i++{
		sum += arr[i]
	} 
	avg = float64(sum) / float64(len(arr))
	return int(avg)
}
func main(){
	var arr [10] int = [10] int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println("average of array is: ", arrAvg(arr))
}