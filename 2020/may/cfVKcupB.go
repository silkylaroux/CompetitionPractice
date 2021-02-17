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
func solution(one int, two int, three int, four int) int {
	//fmt.Println(one, two, three, four)
	if three >= one {
		three -= one
		four += one
		one = 0
	} else {
		one -= three
		four += three
		three = 0
		if two%2 == 1 {
			if one >= 2 {
				one -= 2
			} else {
				one = 0
			}
			two--
			four++
		}
	}
	//fmt.Println(one, two, three, four)
	if two%2 == 1 {
		two--
		four++
	}
	if one%4 > 0 {
		four++
		one -= one % 4
	}
	return four + (three + (two)/2 + (one / 4))
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, ret int
	var one, two, three, four int
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

	for x := 0; x < n; x++ {
		scanner.Scan()
		temp := fastGetInt(scanner.Bytes())
		if temp == 4 {
			four++
		}
		if temp == 3 {
			three++
		}
		if temp == 2 {
			two++
		}
		if temp == 1 {
			one++
		}
	}
	//======================================================
	ret = solution(one, two, three, four)
	//==================== OUTPUT ==========================
	fmt.Println(ret)
}
