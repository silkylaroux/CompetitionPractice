package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
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
		return - n
	}
	return n
}

//========== Implement Algorithm =======================
func solution(b []byte) int {
	if b[0] != '-'{
		return fastGetInt(b)
	}else{
		size := len(b)
		if(b[size-1] < b[size-2]){
			b[size-2] = b[size-1]
		}
		b = b[:size-1]
		return fastGetInt(b)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var ret int
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
	//======================================================
	ret = solution(scanner.Bytes())
	//==================== OUTPUT ==========================
	fmt.Print(ret)
}	
