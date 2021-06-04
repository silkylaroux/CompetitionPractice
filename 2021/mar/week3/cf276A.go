package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
func solution() string {
	return ""
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n int
	var k, fi, ti, maxval, tempMax int64
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
	scanner.Scan()
	k = fastGetInt64(scanner.Bytes())
	//======================================================
	maxval = math.MinInt64
	//==================== OUTPUT ==========================
	for i := 0; i < n; i++ {
		scanner.Scan()
		fi = fastGetInt64(scanner.Bytes())
		scanner.Scan()
		ti = fastGetInt64(scanner.Bytes())
		//fi - (ti - k)

		if ti > k {
			tempMax = fi - (ti - k)
		} else {
			tempMax = fi
		}
		if tempMax > maxval {
			maxval = tempMax
		}
	}
	fmt.Print(maxval)
}
