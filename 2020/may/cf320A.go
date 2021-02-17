package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fastGetInt(b []byte) int64 {
	neg := false
	if b[0] == '-' {
		neg = true
		b = b[1:]
	}
	var n int64
	n = 0
	for _, v := range b {
		n = n*10 + int64(v-'0')
	}
	if neg {
		return -n
	}
	return n
}

//========== Implement Algorithm =======================
func solution(x int64) int64 {
	var temp int64
	temp = x
	var count int64
	count = 0
	for temp != 0 {
		if temp%2 == 1 {
			count++
		}
		temp = temp / 2
	}

	return count
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var x, ret int64
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
