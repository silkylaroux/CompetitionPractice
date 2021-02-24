package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	// "strconv"
)

type usablelist struct {
	endVal	int
	theList	[]int
}

type myList usablelist

func (a myList) addLeft(){
	for i := 0; i < a.endVal; i++ {
		if a.theList[i] == 0 {
			a.theList[i] = 1
			break 
		}
	}
}

func (a myList) addRight(){
	for i := a.endVal-1; i >=0 ; i-- {
		if a.theList[i] == 0 {
			a.theList[i] = 1
			break 
		}
	}
}

func (a myList) remove(val int){
	a.theList[val] = 0
}

func (a myList) print(){
	for i := 0; i < a.endVal; i++ {
		fmt.Print(a.theList[i])
	}
}

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
func solution() string {
	return "hi"
}

func main() {

	scanner := bufio.NewReader(os.Stdin)

	//================== Variables used ====================
	var n int
	var ml myList
	var str string
	// var byteSlice []byte
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
	
	//======================= I/O ==========================
	bl := make([]byte,1)
	fmt.Fscanf(scanner, "%d\n%s", &n, &str)


	l := make([]int,10)
	ml.endVal = 10
	ml.theList = l


   	for _, c := range str {
   		// c:= str[i]

		bl[0] = byte(c)
		if fastGetCharFromByte(bl[0]) == "R"{
			ml.addRight()
		}else if fastGetCharFromByte(bl[0]) == "L"{
			ml.addLeft()
		}else{
			ml.remove(fastGetInt(bl))//fastGetInt([scanner.Bytes()[n]]))
		}
		// ml.print()
		// fmt.Println()
        	
            
            // fmt.Printf("%q [%d]\n", string(c), sz)
        }
    

	ml.print()
	// 
	
	//======================================================

	//==================== OUTPUT ==========================

}
