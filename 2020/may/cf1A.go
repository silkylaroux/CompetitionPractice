package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func fastGetInt(b []byte) int64 {
	neg := false
	if b[0] == '-' {
		neg = true
		b = b[1:]
	}
	var n int64
	for _, v := range b {
		n = n*10 + int64(v-'0')
	}
	if neg {
		return -n
	}
	return n
}

//========== Implement Algorithm =======================
func solution(n int64, m int64, a int64) int64 {
	h := int64(math.Ceil(float64(n) / float64(a)))
	w := int64(math.Ceil(float64(m) / float64(a)))
	return h * w
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, m, a, ret int64
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
	scanner.Scan()
	m = fastGetInt(scanner.Bytes())
	scanner.Scan()
	a = fastGetInt(scanner.Bytes())
	//======================================================

	ret = solution(n, m, a)
	//==================== OUTPUT ==========================
	fmt.Println(ret)
}
