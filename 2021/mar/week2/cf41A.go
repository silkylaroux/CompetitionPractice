package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//========== Implement Algorithm =======================
func solution(str1 []byte, str2 []byte) string {
	if len(str1) != len(str2) {
		return "NO"
	}
	end := len(str1) - 1
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[end] {
			return "NO"
		}
		end--
	}
	return "YES"
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
	firstStr := scanner.Bytes()
	scanner.Scan()
	ret = solution(firstStr, scanner.Bytes())
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Println(ret)
}