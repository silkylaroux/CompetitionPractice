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

func buildIntSlice(b []byte, size int) []int {
	retVal :=make([]int,size)
	curr:=0
	for i := 0; i < len(b); i++ {
		if b[i] == 48 {
			retVal[curr] =0
			curr++
		}else if b[i] == 49{
			retVal[curr] =1
			curr++
		}
	}

	return retVal

}

//========== Implement Algorithm =======================

func solution(a []int, size int) int {

    max_so_far := -1
    zeroCount := 0
 	oneCount := 0

    
    for i := 0; i < size; i++{
	    if (a[i]==1) {
			oneCount++;
			if (zeroCount>0){
				zeroCount--;
			}
		} else {
			zeroCount++;
			if(zeroCount>max_so_far){
				max_so_far=zeroCount;
			}
		}
    }
    return max_so_far + oneCount;

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	var ret int
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
	scanner.Split(bufio.ScanLines)
	//======================= I/O ==========================
	scanner.Scan()

	size:= fastGetInt(scanner.Bytes())

	scanner.Scan()
	ret = solution(buildIntSlice(scanner.Bytes(),size),size)
	//======================================================

	//==================== OUTPUT ==========================
	fmt.Println(ret)
}

// tried to implement S_Adik_K's method