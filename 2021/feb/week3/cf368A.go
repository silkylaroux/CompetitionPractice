package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

func fastGetCharFromByte(b byte) string {
	t := ""
	if b >= 'A' && b <= 'Z' {
		t = string('Z' - ('Z' - b))
	}
	if b >= 'a' && b <= 'z' {
		t = string('z' - ('z' - b))
	}

	return t
}


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
func solution(scanner *bufio.Scanner, n,m int) string {
	var ch string
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			scanner.Scan()
			ch = fastGetCharFromByte(scanner.Bytes()[0])
			if ch != "B" {
				if ch != "W" {
					if ch != "G" {
						return "#Color"
					}
				}
			}
		}
	}
	return "#Black&White"
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var n,m int
	var ch string
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
	ch = solution(scanner,n,m)
	
	//======================================================
	
	//==================== OUTPUT ==========================
	fmt.Println(ch)
}
