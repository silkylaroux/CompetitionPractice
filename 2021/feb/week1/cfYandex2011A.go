
package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

func fastGetInt(b []byte) int64 {
	neg := false
	if b[0] == '-' {
		neg = true
		b = b[1:]
	}
	var n int64
	n = 0
	for _, v := range b {
		n = n*10 + int64(v-'0')
	}
	if neg {
		return - n
	}
	return n
}

// func Doubled(n int)(result int64) {
// 	result = 1
// 	for i := 0; i < n; i++ {
// 		result *=2
// 	}
// 	return result
// }

//========== Implement Algorithm =======================
func solution(val int64) string {
	var factVal int64
	listNames := []string{"Sheldon", "Leonard", "Penny", "Rajesh", "Howard"}

	factVal = 1
	for val > (5 * factVal) {
		val -= (5 * factVal)
		factVal *= 2
	}
	return listNames[(val-1)/factVal]

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var ret string
	var n int64
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
	ret = solution(n)
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
