package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
func solution(x int) int {
	retVal := int(math.Ceil(float64(x) / 5))
	return retVal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var ret, x int
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
	x = fastGetInt(scanner.Bytes())
	//======================================================
	ret = solution(x)
	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
