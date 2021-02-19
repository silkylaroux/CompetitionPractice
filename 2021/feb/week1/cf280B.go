package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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
		return -n
	}
	return n
}

func Max(x, y float64) float64 {
 if x > y {
   return float64(x)
 }
 return float64(y)
}

//========== Implement Algorithm =======================
func solution(strtLen int, values []int) float64 {
	var retVal float64
	sort.Ints(values)

	retVal = Max(float64(values[0]), float64(strtLen - values[len(values)-1]))
	for i := 0; i < len(values)-1; i++ {
		retVal = Max(retVal, float64((values[i+1]- values[i]))/2.0)
	}
	return retVal

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, max int
	var ret float64
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
	values := make([]int, n)
	scanner.Scan()
	max = fastGetInt(scanner.Bytes())

	for i := 0; i < n; i++ {
		scanner.Scan()
		temp := fastGetInt(scanner.Bytes())
		values[i] = temp

	}
	ret = solution( max , values)
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Printf("%f",ret)
}
// Shoutouts to S_Adik_K 