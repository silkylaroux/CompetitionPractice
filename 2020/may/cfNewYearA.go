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
func solution(inList []int, t int, n int) int {
	temp := 1
	//fmt.Println(temp, t)
	if t == temp {
		return temp
	}
	for temp < t && temp < n {
		//fmt.Println(temp, t)
		temp += inList[temp-1]
		if temp == t {
			return t
		}

	}

	//fmt.Println(temp, t)
	return temp

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var t, n, ret int
	var inList []int
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
	t = fastGetInt(scanner.Bytes())

	for x := 0; x < n-2; x++ {
		scanner.Scan()
		newNum := fastGetInt(scanner.Bytes())
		inList = append(inList, newNum)

	}
	inList = append(inList, 1)

	//======================================================
	ret = solution(inList, t, n)
	//==================== OUTPUT ==========================
	if ret == t {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
