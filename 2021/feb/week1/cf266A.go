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
		return -n
	}
	return n
}


func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//========== Implement Algorithm =======================
func solution(n,m,a,b int) int {
	var retVal int
	if a*m > b {
		trips := (n/m)
		retVal = (n/m)*b
		leftOver := n%m
		if leftOver*a < b {
			for leftOver>0 {
				retVal += a
				leftOver--
			}
		}else{
			for n-trips*m >0 {
				trips+=m
				retVal += b
			}
		}
	}else{
		return n * a
	}
	
	
	return retVal

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n,m,a,b, ret int
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
	scanner.Scan()
	m = fastGetInt(scanner.Bytes())
	scanner.Scan()
	a = fastGetInt(scanner.Bytes())
	scanner.Scan()
	b = fastGetInt(scanner.Bytes())

	ret = solution( n,m,a,b)
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
// Shoutouts to S_Adik_K 