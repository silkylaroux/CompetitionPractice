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

//int(math.Max(float64((a+b)*c), float64((a+c)*b)))
//========== Implement Algorithm =======================
func solution(a int, b int, c int) int {
	retVal := a * b * c
	if a == 1 || b == 1 || c == 1 {

		retVal = int(math.Max(float64(retVal), float64((a+b)*c)))
		retVal = int(math.Max(float64(retVal), float64((a * (b + c)))))
		retVal = int(math.Max(float64(retVal), float64(a+b+c)))
	}

	return retVal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var a, b, c, ret int
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
	//======================================================
	ret = solution(a, b, c)
	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
