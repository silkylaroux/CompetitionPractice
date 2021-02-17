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
func solution(x int, y int, z int) bool {

	return x == 0 && y == 0 && z == 0
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var x, y, z, num int
	var ret bool
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
	for i := 0; i < num; i++ {
		scanner.Scan()
		x += fastGetInt(scanner.Bytes())
		scanner.Scan()
		y += fastGetInt(scanner.Bytes())
		scanner.Scan()
		z += fastGetInt(scanner.Bytes())
		//fmt.Println(x, y, z)
	}
	//======================================================

	ret = solution(x, y, z)

	//==================== OUTPUT ==========================
	if ret {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
