package main

import (
	//"bufio"
	"fmt"
)

/*
func solveFromFile() int64 {
	fmt.Println("here")
	var nu int
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nu = fastGetInt(scanner.Bytes())

	defer file.Close()

	return -1 //solution(nu)
}*/

//========== Implement Algorithm =======================
func solution(num1 uint64, num2 uint64) uint64 {
	var retval uint64
	if num2 <= num1/2+num1%2 {
		retval = num2*2 - 1
	} else {
		if num1%2 == 1 {
			retval = num1 - num1%num2*2 - 1
		} else {
			retval = num1 - num1%num2*2
		}

	}
	return retval
}

func main() {

	//================== Variables used ====================
	var num1, num2, ret uint64
	//======================================================

	//if len(os.Args) >= 2 {
	//============== Get return value from File ============
	//ret = solveFromFile()

	//} else {
	//============== Get return value from STDIN ===========
	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	//scanner.Scan()
	//num = fastGetInt(scanner.Bytes())
	ret = solution(num1, num2)
	//	}

	//================== OUTPUT ============================
	fmt.Println(ret)
}