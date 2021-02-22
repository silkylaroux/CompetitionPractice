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
		return - n
	}
	return n
}

//========== Implement Algorithm =======================
func solution(a,b int) int {
	if a == b {
		return 0
	}
	if a%2 == 0{
		if b%2 == 0 {
			if a < b {
				return 2
			}else{
				return 1
			}
		}else{
			if a < b {
				return 1
			}else{
				return 2
			}
		}
	}else{
		if b%2 == 0{
			if a < b {
				return 1
			}else{
				return 2
			}
		}else{
			if a < b {
				return 2
			}else{
				return 1
			}
		}
	}
	return -1
}


func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n int
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
	for i := 0; i < n; i++ {
		scanner.Scan()
		a:= fastGetInt(scanner.Bytes())

		scanner.Scan()
		b:= fastGetInt(scanner.Bytes())

		fmt.Print(solution(a,b),"\n")
	}
	//======================================================

	//==================== OUTPUT ==========================
}
