package main
import "fmt"
func main(){
	var country map[string]string
	country = make(map[string]string)
	country ["India"] = "IN"
	country ["United kingdom"] = "UK"
	country ["France"] = "Fr"
    fmt.Println(country)
	delete(country,"France")
}