package main

import (
	"bufio"
	"log"
	"os"
)

//========== Implement Algorithm =======================
func solution() string {
	return ""
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================

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

	//======================================================

	//==================== OUTPUT ==========================

}
