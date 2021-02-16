package main

import (
	"fmt"
	"math"
  //"bufio"
  //"os"
)

func main() {
  //cin := bufio.NewReader(os.Stdin)
	//cout := bufio.NewWriter(os.Stdout)
	//defer cout.Flush()

    var k,n,w int
    fmt.Scan(&k)
    fmt.Scan(&n)
    fmt.Scan(&w)
    money:=0
	for x := 1; x < w+1; x++ {
    money+= k*x
    //fmt.Println(money)
	}
    
	fmt.Println(int(math.Max(float64(money -n),0)))

	//fmt.Println(strings.Compare(numStops,B))

}