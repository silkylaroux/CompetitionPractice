package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func equalSlice(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for x := 0; x < len(s1); x++ {
		if s1[x] != s2[x] {
			return false
		}
	}
	return true
}

//========== Implement Algorithm =======================
func solution(sl []int) {
	sort.Ints(sl)
	check := false
	for i := 1; i < 1040; i++ {
		newSl := make([]int, len(sl))
		for x := 0; x < len(sl); x++ {
			newSl[x] = sl[x] ^ i
		}
		sort.Ints(newSl)
		if equalSlice(newSl, sl) {
			fmt.Println(i)
			check = true
			break
		}
	}
	if !check {
		fmt.Println(-1)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var num int
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
	//======================================================
	for x := 0; x < num; x++ {
		scanner.Scan()
		a := fastGetInt(scanner.Bytes())
		sl := make([]int, a)
		for i := 0; i < a; i++ {
			scanner.Scan()
			sl[i] = fastGetInt(scanner.Bytes())
		}
		solution(sl)
	}
	//==================== OUTPUT ==========================

}