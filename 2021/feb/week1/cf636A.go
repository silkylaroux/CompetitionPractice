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
		return -n
	}
	return n
}


//========== Implement Algorithm =======================
func solution(val int) int {
	for i := 3; i <= val; i = 1 + i*2 {
		if val % i == 0{
			return val / i
		}
	}
	return -1
	
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var t, ret, in int
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
	t = fastGetInt(scanner.Bytes())

	for i := 0; i < t; i++ {
		scanner.Scan()
		in = fastGetInt(scanner.Bytes())
		ret=solution(in)
		fmt.Print(ret)
		fmt.Print("\n")
	}
	//======================================================

	//==================== OUTPUT ==========================
}
