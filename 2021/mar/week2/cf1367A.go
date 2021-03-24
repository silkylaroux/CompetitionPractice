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
func solution(val string) string {
	retVal := val[0:2]
	//fmt.Println(retVal)
	for x := 3; x < len(val); x = x + 2 {
		retVal += string(val[x])
	}
	return retVal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var num int
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
	num = fastGetInt(scanner.Bytes())
	for x := 0; x < num; x++ {
		scanner.Scan()
		ret = solution(scanner.Text())
		fmt.Println(ret)
	}
	//======================================================

	//==================== OUTPUT ==========================

}
