package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//========== Implement Algorithm =======================
func solution(str string) bool {
	if len(str) < 5 {
		return false
	}
	h := false
	e := false
	l := false
	o := false
	lcount := 0
	for x := 0; x < len(str); x++ {
		if h && e && l && o {
			return true
		}

		if string(str[x]) == "h" {
			h = true
		} else if string(str[x]) == "e" && h {
			e = true
		} else if string(str[x]) == "l" && h && e {
			lcount++
			if lcount == 2 {
				l = true
			}

		} else if string(str[x]) == "o" && h && e && l {
			o = true
		}
	}

	return h && e && l && o
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var str string
	var ret bool
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
	str = scanner.Text()
	//======================================================
	ret = solution(str)
	//==================== OUTPUT ==========================
	if ret {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
