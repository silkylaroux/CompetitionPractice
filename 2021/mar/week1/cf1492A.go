package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"math"
)

func fastGetInt64(b []byte) int64 {
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
		return -n
	}
	return n
}
 
func getDiff(a,p int64) int64{
	return a - (p%a)
}
 
func getMin(a,b int64) int64{
 
	if a < b{
		return a
	}
	return b
}
 
//========== Implement Algorithm =======================
func solution(p,a,b,c int64) int64 {
	
	if p%a == 0 || p%b == 0 || p%c == 0 {
		return 0
	}
 
	val := getMin(getDiff(a,p),math.MaxInt64)
	val = getMin(getDiff(b,p),val)
	val = getMin(getDiff(c,p),val)
 
	return val
 
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, p, a, b, c  int64
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
		n = fastGetInt64(scanner.Bytes())
		for i := 0; i < int(n); i++ {
			scanner.Scan()
			p = fastGetInt64(scanner.Bytes())
			scanner.Scan()
			a = fastGetInt64(scanner.Bytes())
			scanner.Scan()
			b = fastGetInt64(scanner.Bytes())
			scanner.Scan()
			c = fastGetInt64(scanner.Bytes())
			fmt.Println(solution(p,a,b,c))
		}
	//======================================================
	
	//==================== OUTPUT ==========================

}
