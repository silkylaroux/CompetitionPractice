package main
 
import (
	"fmt"
)
 
func main() {
	var numStops, total, final int
    _, err1 := fmt.Scanf("%d",&numStops)
    //.Println(numStops)
	for x := 0; x < numStops; x++ {
        var out, in int
        fmt.Scan(&out)
        fmt.Scan(&in)
        //fmt.Println(out,in)
    //temp = append(temp, out,in)
    
		total += in-out
		//.Println(total)
		if final < total{
			final = total
		}
 
 
	}
	if err1!= nil{
	    panic(err1)
	}
	fmt.Println(final)
 
	//fmt.Println(strings.Compare(numStops,B))
 
}