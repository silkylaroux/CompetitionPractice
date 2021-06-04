package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		return -n
	}
	return n
}

//========== Implement Algorithm =======================
func solution(n, k int) {
	fmt.Println(n - (k+1)/2)
	for z := (k + 1) / 2; z <= n; z++ {
		if z != k {
			fmt.Printf("%d ", z)
		}
	}
	fmt.Println()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var i, n, k int
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
	i = fastGetInt(scanner.Bytes())
	for z := 0; z < i; z++ {
		scanner.Scan()
		n = fastGetInt(scanner.Bytes())
		scanner.Scan()
		k = fastGetInt(scanner.Bytes())
		solution(n, k)
	}
	//======================================================

	//==================== OUTPUT ==========================

}
