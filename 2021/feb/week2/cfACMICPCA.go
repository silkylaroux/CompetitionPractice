package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"sort"
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
func solution(ints []int, total int) string {
	if total % 2 != 0{
		return "NO"
	}
	half := total/2
	lastIndex := len(ints)-1

	sort.Ints(ints)
	if (ints[lastIndex] + ints[0] + ints[1]) == half || (ints[lastIndex] + ints[2] + ints[1]) == half{
		return "YES"
	}
	if (ints[lastIndex] + ints[lastIndex-1] + ints[0]) == half {
		return "YES"
	}
	if (ints[lastIndex] + ints[lastIndex-1] + ints[lastIndex-2]) == half {
		return "YES"
	}

	for i := 0; i < 6; i++ {
		for x := i+1; x <6; x++ {
			for y := x+1; y < 6; y++ {
				if (ints[i]+ints[x]+ints[y])==half {
					return "YES"
				}
			}
		}
	}
	return "NO"
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, total int
	ints := make([]int, 6)
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
	n = 6
	for i := 0; i < n; i++ {
		scanner.Scan()
		temp := fastGetInt(scanner.Bytes())
		total += temp
		ints[i] = temp
	}

	ret = solution(ints,total)

	//======================================================

	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
