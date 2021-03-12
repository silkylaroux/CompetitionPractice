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
func solution(rows int, cols int) string {
	retVal := "Vivek"
	//fmt.Println(rows, cols)
	if rows > cols {
		rows, cols = cols, rows
	}

	if rows%2 == 1 {
		retVal = "Ashish"

	}
	return retVal
}

func checkNumZero(freeCols []int) int {
	count := 0
	for x := 0; x < len(freeCols); x++ {
		if freeCols[x] == 0 {
			count++
		}
	}
	return count
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var num, x, y, freeRows, freeCols int
	var ret string
	//var freeCol []int
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
	for i := 0; i < num; i++ {
		scanner.Scan()
		x = fastGetInt(scanner.Bytes())
		scanner.Scan()
		y = fastGetInt(scanner.Bytes())
		freeCol := make([]int, y)
		freeRows = 0
		for x1 := 0; x1 < x; x1++ {
			hasFreeRow := true
			for c := 0; c < y; c++ {
				scanner.Scan()
				newNum := fastGetInt(scanner.Bytes())
				if newNum != 0 {
					hasFreeRow = false
					freeCol[c] = 1
				}
			}
			if hasFreeRow {
				freeRows++
			}
		}
		freeCols = checkNumZero(freeCol)
		ret = solution(freeRows, freeCols)
		fmt.Println(ret)
	}
	//======================================================
	//ret = solution(freeRows, freeCols)
	//==================== OUTPUT ==========================
	//fmt.Println(ret)
}
