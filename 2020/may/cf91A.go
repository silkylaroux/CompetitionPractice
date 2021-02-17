package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
func solution(num int) bool {
	retVal := false
	if num%4 == 0 || num%7 == 0 || num%47 == 0 || num%74 == 0 {
		retVal = true
	} else {
		slice := strconv.Itoa(num)
		for i := 0; i < len(slice); i++ {
			if fastGetInt([]byte{slice[i]}) != 4 && fastGetInt([]byte{slice[i]}) != 7 {
				retVal = false
				break
			}
			if i == len(slice)-1 {
				retVal = true
			}
		}
	}

	return retVal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var num int
	var ret bool
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
	//======================================================
	ret = solution(num)
	//==================== OUTPUT ==========================
	if ret {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
