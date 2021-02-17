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

func solveFromFile() []int {
	var num int
	var numList []int
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	num = fastGetInt(scanner.Bytes())
	num++
	for scanner.Scan() {
		numList = append(numList, fastGetInt(scanner.Bytes()))
	}

	defer file.Close()

	return solution(numList)
}

//========== Implement Algorithm =======================
func solution(numList []int) []int {
	sort.Ints(numList)

	return numList
}

func main() {

	//================== Variables used ====================
	var num int
	var ret []int
	//======================================================

	if len(os.Args) >= 2 {
		//============== Get return value from File ============
		ret = solveFromFile()

	} else {
		//============== Get return value from STDIN ===========
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)

		scanner.Scan()
		num = fastGetInt(scanner.Bytes())
		numList := make([]int, num)

		for x := 0; x < num; x++ {
			scanner.Scan()
			numList[x] = fastGetInt(scanner.Bytes())
		}

		ret = solution(numList)
	}
	//================== OUTPUT ============================
	for x := 0; x < len(ret); x++ {
		fmt.Print(ret[x])
		fmt.Print(" ")
	}
}
