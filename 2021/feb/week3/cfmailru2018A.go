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
		return - n
	}
	return n
}

//========== Implement Algorithm =======================
func solution(location int, forward,reverse []int) string {
	location--

	if forward[0] == 0 || (forward[location] == 0 && reverse[location] == 0){
		return "NO"
	}

	if forward[location]==1{
		return "YES"
	}
	if reverse[location] == 1{
		for i := location; i < len(reverse); i++ {
						if (forward[i] == 1 && reverse[i] == 1){
				return "YES"
			}
		}
	}
	return "NO"
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n, b int
	var forward, reverse []int
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
	scanner.Scan()
	n = fastGetInt(scanner.Bytes())
	scanner.Scan()
	b = fastGetInt(scanner.Bytes())
	forward,reverse = make([]int, n), make([]int,n)
	for i := 0; i < n; i++ {
		
		scanner.Scan()
		temp := fastGetInt(scanner.Bytes())
		forward[i]= temp

	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		temp := fastGetInt(scanner.Bytes())
		reverse[i] = temp
	}

	ret = solution(b, forward,reverse)


	
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Println(ret)
}
