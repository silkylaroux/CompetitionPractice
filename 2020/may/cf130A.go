package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	num = fastGetInt(scanner.Bytes())

	defer file.Close()

	return solution(num)
}

//========== Implement Algorithm =======================
func solution(num int) int {
	return num
}

func main() {

	//================== Variables used ====================
	var num, ret int
	//======================================================

	if len(os.Args) >= 2 {
		//============== Get return value from File ============
		ret = solveFromFile()

	} else {
		//============== Get return value from STDIN ===========
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)

		scanner.Scan()
		s := scanner.Text()
		s = strings.ReplaceAll(s, "WUB", " ")
		//num = fastGetInt(scanner.Bytes())
    s = strings.TrimSpace(s)
    slist := strings.Split(s, " ")

		num = 0
    num++
		//fmt.Println(slist)
    for x:=0; x<len(slist); x++{
      if slist[x] != " " && slist[x] != ""{
        fmt.Print(slist[x]+" ")
      }
      
    }
		ret = 0
    ret++
		//ret = solution(num)
	}
	//================== OUTPUT ============================

}
