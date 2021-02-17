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

func solveFromFile() int {
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
	scanner.Scan()
	for scanner.Scan() {
		numList = append(numList, fastGetInt(scanner.Bytes()))
	}

	defer file.Close()

	return solution(num, numList)
}

//========== Implement Algorithm =======================
func solution(num int, numList []int) int {
	minNum := 99999
	sort.Ints(numList)
	if len(numList) == num {
		return numList[num-1] - numList[0]
	}
	for x := 0; x < (len(numList) - (num - 1)); x++ {
		if (numList[x+num-1] - numList[x]) < minNum {
			minNum = (numList[x+num-1] - numList[x])
		}
	}
	return minNum
}

func main() {

	//================== Variables used ====================
	var num, ret int
	var numList []int
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
		scanner.Scan()

		for scanner.Scan() {
			numList = append(numList, fastGetInt(scanner.Bytes()))
		}

		ret = solution(num, numList)
	}

	//================== OUTPUT ============================
	fmt.Println(ret)
}
