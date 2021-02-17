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

type dragon struct {
	strength int
	bonus    int
}

type ByStrength []dragon

func (a ByStrength) Len() int           { return len(a) }
func (a ByStrength) Less(i, j int) bool { return a[i].strength < a[j].strength }
func (a ByStrength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

//========== Implement Algorithm =======================
func solution(str int, drags []dragon) bool {
	sort.Sort(ByStrength(drags))
	for x := 0; x < len(drags); x++ {
		if x < len(drags) && drags[x].strength >= str {
			return false
		}
		str += drags[x].bonus
	}
	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var str, num int
	var ret bool
	var drags []dragon
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
	str = fastGetInt(scanner.Bytes())
	scanner.Scan()
	num = fastGetInt(scanner.Bytes())

	for x := 0; x < num; x++ {
		scanner.Scan()
		s := fastGetInt(scanner.Bytes())
		scanner.Scan()
		bonus := fastGetInt(scanner.Bytes())
		drags = append(drags, dragon{s, bonus})
	}

	//======================================================
	ret = solution(str, drags)
	//==================== OUTPUT ==========================
	if ret {
		fmt.Print("YES")
	} else {
		fmt.Println("NO")
	}
}
