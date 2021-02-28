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

//========== Implement Algorithm =======================
func solution(a,b,c,d, n int ) int {
	count := 0
	for i := 0; i <= n; i++ {
		if i%a == 0 || i % b == 0 || i % c == 0 || i % d == 0 {
		}else{
			count++ 
		}
	}
	return n-count
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var a,b,c,d, n int 
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
	a = fastGetInt(scanner.Bytes())
	scanner.Scan()
	b = fastGetInt(scanner.Bytes())
	scanner.Scan()
	c = fastGetInt(scanner.Bytes())
	scanner.Scan()
	d = fastGetInt(scanner.Bytes())
	scanner.Scan()
	n = fastGetInt(scanner.Bytes())
	fmt.Println(solution(a,b,c,d,n))

	//======================================================
	
	//==================== OUTPUT ==========================

}
