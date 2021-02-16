package main
 
import (
	"fmt"
	"log"
    "strings"
    "sort"
)
 
func main() {
	A:=""
  _, err := fmt.Scanf("%s", &A)
  b:=strings.Split(A,"+")
 
 
	if err !=nil{
		log.Fatal(err)
  }
  sort.Strings(b)
  A=b[0]
  for x:=1; x<len(b);x++{
      A += "+" + b[x]
  }
  fmt.Print(A)
  
}