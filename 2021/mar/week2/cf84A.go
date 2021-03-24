package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//========== Implement Algorithm =======================
func solution(byt []byte) string {
	count := 0
	for x := 0; x < len(byt); x++ {
		if byt[x] == '4' || byt[x] == '7' {
			count++
		}
	}
	if count == 4 || count == 7 {
		return "YES"
	}
	return "NO"
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
