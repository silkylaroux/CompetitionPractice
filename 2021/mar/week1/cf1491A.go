package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)
// sortable struct      sort.Sort(ByStrength(drags))
type dragon struct {
	strength int
	bonus    int
}
// ByStrength returns 
type ByStrength []dragon

func (a ByStrength) Len() int           { return len(a) }
func (a ByStrength) Less(i, j int) bool { return a[i].strength < a[j].strength }
func (a ByStrength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

//========== Implement Algorithm =======================
func solution(t,k int, list []dragon, oneCount int) int {
	if t != 2 {
		val := list[k-1].strength
		p := &list[k-1]
		p.strength = 1-val
		if val == 1{
			oneCount--
		}else{
			oneCount++
		}
	}else{
		if oneCount >= k {
			fmt.Println(1)
		}else{
			fmt.Println(0)
		}
	}
	return oneCount
	

}

func main() {

	scanner := bufio.NewReader(os.Stdin)

	//================== Variables used ====================
	var n,q,t,k int
	
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
	fmt.Fscan(scanner, &n, &q)
	list := make([]dragon,n)
	// sortedList := make([]dragon,n)
	var in,oneCount int
	for i := 0; i < n; i++ {

		fmt.Fscan(scanner, &in)
		// scanner.Scan()
		// inval := fastGetInt(scanner.Bytes())
		val := dragon{in,0}
		if in == 1{
			oneCount++
		}
		list[i] = val
		// sortedList[i] = val
	}

	// sort.Sort(ByStrength(sortedList))
		
	for i := 0; i < q; i++ {
		fmt.Fscan(scanner, &t, &k)
		oneCount = solution(t,k,list,oneCount)
	}


	//======================================================
	
	//==================== OUTPUT ==========================

}
