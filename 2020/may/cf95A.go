package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//========== Implement Algorithm =======================
func solution(str string) string {
	alpha := "qwertyuiopasdfghjklzxcvbnm"

	if !strings.ContainsAny(str, alpha) {
		str = strings.ToLower(str)
	} else if !strings.ContainsAny(string(str[1:]), alpha) {
		strTemp := strings.ToUpper(string(str[0]))
		str = strTemp + strings.ToLower(string(str[1:]))
	}
	return str
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var str, ret string
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
	fmt.Println(ret)
}
