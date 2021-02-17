package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//========== Implement Algorithm =======================
func solution(a int, b int) int {
	retVal := 0
	for a <= b {
		a *= 3
		b *= 2
		retVal++
	}
	return retVal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var a, b, ret int
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
	fmt.Scan(&a, &b)
	//======================================================
	ret = solution(a, b)
	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
