package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isUpperCase(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

//========== Implement Algorithm =======================
func solution(str []byte) string {
	upperCount := 0
	lowerCount := 0
	for x := 0; x < len(str); x++ {
		if isUpperCase(str[x]) {
			upperCount++
		} else {
			lowerCount++
		}
	}
	if upperCount > len(str)/2 {
		return strings.ToUpper(string(str))
	}
	if lowerCount >= len(str)/2 {
		return strings.ToLower(string(str))
	}
	return "fake"
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
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
	ret = solution(scanner.Bytes())
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Println(ret)
}
