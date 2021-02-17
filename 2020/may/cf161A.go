package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//========== Implement Algorithm =======================
func solution(a int, b int) int {
	retVal := 0
	for a <= b {
		a *= 3
		b *= 2
		retVal++
	}
	return retVal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var a, b, c, d, e, ret int
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

	for i := 1; i < 6; i++ {
	    if i ==1 || i ==5{ret = 2}
	    if i ==2 || i ==4{ret = 1}
        if i == 3 { ret =0}
		fmt.Scan(&a, &b, &c, &d, &e)
		if a != 0 || e != 0 {
			ret += 2
			break
		} else if b != 0 || d != 0 {
			ret++
			break
		} else if c != 0 {
			break
		}
	}

	//======================================================
	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
