package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strconv"
)

func fastGetInt(b []byte) int {
	neg := false
	if b[0] == '-' {
		neg = true
		b = b[1:]
	}
	var n int
	n = 0
	for _, v := range b {
		n = n*10 + int(v-'0')
	}
	if neg {
		return - n
	}
	return n
}

func max(a,b int) int{
	if a > b{
		return a 
	}
	return b
}

func min(a,b int) int{
	if a < b{
		return a 
	}
	return b
}

func dist(a,b int) int{
	
	return abs(a-b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toStr(a int) string{
	return strconv.Itoa(a)
}



//========== Implement Algorithm =======================
func solution(list []int)  {
	front := list[0]
	end := list[len(list)-1]
	// retval := ""

	// retval += toStr(dist(front,list[1])) + " " + toStr(dist(front,end))+"\n"
	fmt.Printf("%d %d\n", dist(front,list[1]), dist(front,end))
	var a,b,minVal,maxVal int
	for i := 1; i < len(list)-1; i++ {
		
		a = dist(list[i],list[i-1])
		b = dist(list[i],list[i+1])
		minVal = min(a,b)
		maxVal = max(dist(list[i],front),dist(list[i],end))
		fmt.Printf("%d %d\n", minVal, maxVal)
		// retval += toStr(minVal) + " "+ toStr(maxVal) +"\n"
	}
	fmt.Printf("%d %d", dist(list[len(list)-2],list[len(list)-1]), dist(front,list[len(list)-1]))
	// retval += toStr(dist(list[len(list)-2],list[len(list)-1])) + " " + toStr(dist(front,list[len(list)-1]))

	// return retval
}

func main() {

	scanner := bufio.NewReader(os.Stdin)


	

	//================== Variables used ====================
	var n int
	// var ret string
	//======================================================

	//============== Get return value from File ============
	if len(os.Args) >= 2 {

		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		scanner = bufio.NewReader(file)
		defer file.Close()

	}

	
	// scanner.Split(bufio.ScanWords)
	//======================= I/O ==========================
	// scanner.Scan()
	// n = fastGetInt(scanner.Bytes())
	// list := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	scanner.Scan()
	// 	list[i]=fastGetInt(scanner.Bytes())
	// }
	fmt.Fscanf(scanner, "%d\n", &n)
	list := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(scanner, "%d ", &list[i])
	}
	solution(list)
	//======================================================

	//==================== OUTPUT ==========================
	// fmt.Print(ret)
}
