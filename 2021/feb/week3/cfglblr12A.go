package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"sort"
)

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

func fastGetStrFromBytes(b []byte) string {
	t := ""
	for x := 0; x < len(b); x++ {
		t = t + fastGetCharFromByte(b[x])
	}
	return t
}

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
func solution(val []byte) string {
	
	sort.Slice(val, func(i int, j int) bool { return val[i] < val[j] })
	return fastGetStrFromBytes(val)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n int
	var ret string
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
		scanner.Scan()
		val := scanner.Bytes()

		ret = solution(val)
		fmt.Println(ret)
	}
	//======================================================

	//==================== OUTPUT ==========================

}
