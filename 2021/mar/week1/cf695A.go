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
func solution(val int) string {
	retVal := "9"
	strVal := "0123456789"
	if val == 1 {
		return retVal
	}
	if val == 2 {
		retVal +="8"
		return retVal
	}
	if val == 3 {
		retVal +="89"
		return retVal
	}
	retVal +="89"
	val-=3
	for val >=10 {
		retVal+=strVal
		val -=10
	}
	temp := (val)
	if temp < 0 {
		temp = 0 
	}
	// fmt.Println("here",strVal[:temp])
	retVal+=strVal[:temp]
	
	return retVal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, val int
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
		val = fastGetInt(scanner.Bytes())
		fmt.Println(solution(val))
	}
	//======================================================
	
	//==================== OUTPUT ==========================

}
