package main
 
import (
	"fmt"
    "strings"
)
 
func main() {
	var A string
	var B string
  _, err := fmt.Scan(&A)
  _, err1 := fmt.Scan(&B)
  A = strings.ToLower(A)
  B = strings.ToLower(B)
  
  if err != nil||err1 != nil{
      fmt.Println("blah")
  }
  
  fmt.Println(strings.Compare(A,B))
  
}