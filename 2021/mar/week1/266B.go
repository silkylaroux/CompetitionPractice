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

func fastGetInt64(b []byte) int64 {
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
func solution(a int64, b int64) {
	count := 0
	if a == b {
		count = 0
	} else if a%2 == 1 && b%2 == 1 {
		count = -1
	} else if a > b {
		for a > b {
			if a%8 == 0 && a/8 >= b {
				a /= 8
				count++
			} else if a%4 == 0 && a/4 >= b {
				a /= 4
				count++
			} else if a%2 == 0 && a/2 >= b {
				a /= 2
				count++
			} else if count == 0 {
				count = -1
				break
			} else {
				break
			}
		}
	} else {
		for a < b {
			if a*8 <= b {
				a *= 8
				count++
			} else if a*4 <= b {
				a *= 4
				count++
			} else if a*2 <= b {
				a *= 2
				count++
			} else if count == 0 {
				count = -1
				break
			} else {
				break
			}
		}
	}
	if a != b {
		count = -1
	}
	fmt.Println(count)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var num int
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
	num = fastGetInt(scanner.Bytes())
	for x := 0; x < num; x++ {
		scanner.Scan()
		a := fastGetInt64(scanner.Bytes())
		scanner.Scan()
		b := fastGetInt64(scanner.Bytes())

		solution(a, b)
	}
	//======================================================

	//==================== OUTPUT ==========================

}