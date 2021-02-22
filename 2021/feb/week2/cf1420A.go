package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"sort"
)

type Stack []int

func (s *Stack) addToStack(str int){
	*s = append(*s, str)
}

func (s *Stack) sortStack(){
	sort.Slice(*s, func(p, q int) bool {  
     return (*s)[p] < (*s)[q] })
}

func (array *Stack) insertInt( value int, index int) {
	*array = append((*array)[:index], append([]int{value}, (*array)[index:]...)...)
}

func (array *Stack) removeInt( index int) {
	*array = append((*array)[:index], (*array)[index+1:]...)
}

func (array *Stack) moveInt(srcIndex int) {
    value := (*array)[srcIndex]
    array.removeInt(srcIndex)
    array.insertInt(value, len(*array))
}

func (s *Stack) indexOf(str, maxLen int) int{
    
	for i := 0; i < maxLen; i++ {
		if (*s)[i] == str {
			return i
		}
	}
	return maxLen
}

func (s *Stack) IsSorted() bool {

	for i := 0; i < len(*s)-1; i++ {
		if (*s)[i] <= (*s)[i+1]{
			return true
		}
	}
	return false
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
func solution(list Stack, tempLen int) string {
    
    isOk := list.IsSorted()
	
    if isOk {
    	return "YES"
    }
    return "NO"

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//================== Variables used ====================
	
	var n,tempLen int
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
	for i := 0; i < n; i++ {

		scanner.Scan()
		tempLen = fastGetInt(scanner.Bytes())
		var list Stack
		for x := 0; x < tempLen; x++ {
			scanner.Scan()
			list.addToStack(fastGetInt(scanner.Bytes()))
		}
		ret += solution(list,tempLen) + "\n"
	}

	//======================================================

	//==================== OUTPUT ==========================
	fmt.Print(ret)
}
