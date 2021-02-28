package main

import (
	"bufio"
	"log"
	"os"
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

func isEven(a int) bool{
	return a%2==0
}

func doMath(a,b int) int{
	return (a/2) * b 
}

//========== Implement Algorithm =======================
func solution(a,b int) int {

	if isEven(a) {
		return doMath(a,b)
	}

	if isEven(b){
		return doMath(b,a)
	}

	return doMath(a,b) + b/2+1
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, a, b int
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
		a = fastGetInt(scanner.Bytes())
		scanner.Scan()
		b = fastGetInt(scanner.Bytes())
		fmt.Println(solution(a,b))
	}
	//======================================================

	//==================== OUTPUT ==========================

}
