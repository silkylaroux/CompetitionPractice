package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func fastGetCharFromByte(b byte) string {
	t := ""
	if b >= 'A' && b <= 'Z' {
		t = string('Z' - ('Z' - b))
	}
	if b >= 'a' && b <= 'z' {
		t = string('z' - ('z' - b))
	}

	return t
}

//========== Implement Algorithm =======================
func solution(lineStr string, seconds int) string {
	// make a slice of chars
	lineSlice := strings.Split(lineStr, "")
	//fmt.Println(lineSlice)
	numSwaps := 0

	for x := 0; x < seconds; x++ {
		numSwaps = 0
		prev := lineSlice[0]
		for i := 1; i < len(lineSlice); i++ {
			curr := lineSlice[i]
			if prev == "B" && curr == "G" {
				numSwaps++
				lineSlice[i] = "B"
				lineSlice[i-1] = "G"

			}

			//fmt.Println("prev:", prev, "curr:", curr)
			prev = curr
		}
		//fmt.Println("line:", lineSlice)
		if numSwaps == 0 {
			return strings.Join(lineSlice, "")
		}
	}
	//fmt.Println(lineSlice)
	return strings.Join(lineSlice, "")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var lineStr, ret string
	var seconds int
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
	scanner.Scan()
	seconds = fastGetInt(scanner.Bytes())
	scanner.Scan()
	lineStr = scanner.Text()
	//======================================================
	ret = solution(lineStr, seconds)
	//==================== OUTPUT ==========================
	fmt.Println(ret)
}