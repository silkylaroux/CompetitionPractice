package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"sort"
	"strconv"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str int) {
	
	if len(*s) == 0  {
		*s = append(*s, str) 
	}else{
		*s = append(*s, 0)
		i := sort.Search(len(*s)-1, func(i int) bool { return (*s)[i] > str })

		copy((*s)[i+1:], (*s)[i:])
		(*s)[i] = str
	}
	
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1 
		element := (*s)[index] 
		*s = (*s)[:index] 
		return element, true
	}
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
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
func solution(scanner *bufio.Reader, n int)  {
	m := map[int]int{}
	myInt :=0
	put := n
	// var ok bool
	for i := 0; i < n; i++ {
		// scanner.Scan()
		
		fmt.Fscan(scanner, &myInt)
		m[myInt] = myInt

		for m[put]!=0{
			fmt.Print(strconv.Itoa(put)+" ");
			put--;
		}
		if i != n-1{
				fmt.Println("")
		}
	}
		// list[i] = fastGetInt(scanner.Bytes())
		
	// }
	// var retVal string
	// // sortedList := make([]int, len(list))
	// // copy(sortedList, list)
	// var stack Stack
	// currPlace := size
	// // sort.Ints(sortedList)
	// for i := 0; i < len(list); i++ {
	// 	stack.Push(list[i])
		
	// 		tempCount := 0 
	// 		popVal,tempBool := stack.Pop()
	// 	if(list[i] <= popVal){
	// 		for tempBool {
	// 			if (currPlace <= popVal) && tempCount != i+1{
	// 				currPlace--
	// 				retVal += strconv.Itoa(popVal)+ " "
	// 				tempCount++

	// 			}else{
	// 				stack.Push(popVal)
	// 				break
	// 			}

	// 			popVal,tempBool = stack.Pop()
	// 		}
	// 	}else{
	// 		stack.Push(popVal)
	// 	}
	// 		if i != len(list)-1{
	// 			retVal += "\n"
	// 		}
		
	// }
	

	// return retVal
}

func main() {

	scanner := bufio.NewReader(os.Stdin)

	//================== Variables used ====================
	var n int
	// var ret string

	//======================================================

	//============== Get return value from File ============
	if len(os.Args) >= 2 {

		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		scanner = bufio.NewReader(file)
		defer file.Close()

	}
	// scanner.Split(bufio.ScanWords)
	//======================= I/O ==========================
	fmt.Fscan(scanner, &n)
	// var m map[int]int
	// n = fastGetInt(scanner.Bytes())
	
	solution(scanner,n)
	//======================================================

	//==================== OUTPUT ==========================
	// fmt.Print(ret)
}
