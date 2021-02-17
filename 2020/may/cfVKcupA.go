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
func solution(inList []int, max int) string {
	return ""
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var total, kth, count int

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
	total = fastGetInt(scanner.Bytes())
	scanner.Scan()
	kth = fastGetInt(scanner.Bytes())
	//inList := make([]int, total)
	count = 0
	temp := 0
	for temp < kth {
		scanner.Scan()
		if fastGetInt(scanner.Bytes()) > 0 {
			count++
		}
		temp++
	}
	currVal := fastGetInt(scanner.Bytes())
	kth = currVal
	for temp < total {
		scanner.Scan()
		currVal = fastGetInt(scanner.Bytes())
		if currVal < kth {
			break
		} else if currVal > 0 {
			count++
		}
		temp++
	}
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Println(count)
}
