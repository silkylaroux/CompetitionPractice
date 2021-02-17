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

func solveFromFile() int {
	var n, m int
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n = fastGetInt(scanner.Bytes())
	scanner.Scan()
	m = fastGetInt(scanner.Bytes())

	defer file.Close()

	return solution(n, m)
}

//========== Implement Algorithm =======================
func solution(n int, m int) int {
	small := int(math.Min(float64(n), float64(m)))
	if (small)%2 == 0 {
		return 0
	}
	return 1
}

func main() {

	//================== Variables used ====================
	var n, m, ret int
	//======================================================

	if len(os.Args) >= 2 {
		//============== Get return value from File ============
		ret = solveFromFile()

	} else {
		//============== Get return value from STDIN ===========
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)

		scanner.Scan()
		n = fastGetInt(scanner.Bytes())
		scanner.Scan()
		m = fastGetInt(scanner.Bytes())

		ret = solution(n, m)
	}

	//================== OUTPUT ============================
	if ret == 0 {
		fmt.Println("Malvika")
	} else {
		fmt.Println("Akshat")
	}
}
