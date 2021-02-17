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
	n := 0
	for _, v := range b {
		n = n*10 + int(v-'0')
	}
	if neg {
		return -n
	}
	return n
}

//========== Implement Algorithm =======================
func solution(n int, m int) int {
	count := 0
	temp := 0
	for x := n; x > 0; x-- {
		temp++
		if temp == m {
			x++
			temp = 0
		}
		count++
	}
	return count
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, m, ret int
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
	//======================= I/O ==========================
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n = fastGetInt(scanner.Bytes())
	scanner.Scan()
	m = fastGetInt(scanner.Bytes())
	//======================================================
	ret = solution(n, m)

	//==================== OUTPUT ==========================
	fmt.Println(ret)
}
