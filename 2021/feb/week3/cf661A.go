package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"fmt"
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

//========== Implement Algorithm =======================
func solution(list []int) string {
	var currVal int
	sort.Ints(list)
	hadJump := false
	prevVal := list[0]
	for i := 1; i < len(list); i++ {
		currVal = list[i]
		if (currVal - prevVal) >1 || (currVal - prevVal) < -1 {
				hadJump = true
		}
		prevVal = currVal
	}

	
	
	if hadJump{
		return "NO"
	}else{
		return "YES"
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, currVal int
	//======================================================

	//============== Get return value from File ============
	if len(os.Args) >= 2 {

		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		scanner = bufio.NewScanner(file)
		defer file.Close()

	}
	scanner.Split(bufio.ScanWords)
	//======================= I/O ==========================
	scanner.Scan()
	n = fastGetInt(scanner.Bytes())
	for i := 0; i < n; i++ {
		scanner.Scan()
		currVal = fastGetInt(scanner.Bytes())
		list := make([]int,currVal)
		
		for j := 0; j < currVal; j++ {
			scanner.Scan()
			list[j] = fastGetInt(scanner.Bytes())
			
		}
		fmt.Println(solution(list))
		
		
	}
	//======================================================

	//==================== OUTPUT ==========================

}
