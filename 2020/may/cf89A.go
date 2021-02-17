package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func fastGetStrFromBytes(b []byte) string {
	t := ""
	for x := 0; x < len(b); x++ {
		t = t + fastGetCharFromByte(b[x])
	}
	return t
}

//========== Implement Algorithm =======================
func solution(s string) string {
	str := strings.ToLower(s)
	vowels := "aeiouy"

	for x := 0; x < len(vowels); x++ {
		str = strings.ReplaceAll(str, string(vowels[x]), "")
	}
	return str
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
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
	//======================================================
	ret = solution(fastGetStrFromBytes(scanner.Bytes()))
	//==================== OUTPUT ==========================
	for x := 0; x < len(ret); x++ {
		fmt.Print("." + string(ret[x]))
	}
}
