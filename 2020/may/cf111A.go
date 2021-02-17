package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  "strings"
  "sort"
)

func fastGetInt(b[]byte) int{
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

func solveFromFile() int{
    file, err := os.Open(os.Args[1])
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    scanner.Scan()
  //  fmt.Println(byts)
    

    defer file.Close()

    return solution(strings.Fields(scanner.Text()))
}

//========== Implement Algorithm =======================
func solution(str []string) int{
  var num []int
  count :=0
  for x := 0; x < len(str); x++ {
    i := fastGetInt([]byte(str[x]))
    num = append(num,i)
    count +=i
  }
  
  tempCount := 0
  totNum := 0
  sort.Ints(num)
  for x := len(num)-1; x >= 0; x-- {
    if tempCount<= (count/2){
      tempCount+= num[x]
      totNum++
    }
  }
  return totNum
}

func main() {

//================== Variables used ====================
  var ret int
//======================================================

  if len(os.Args) >=2{
//============== Get return value from File ============
    ret = solveFromFile()


  }else{
//============== Get return value from STDIN ===========
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    scanner.Scan()
    ret = solution(strings.Fields(scanner.Text()))
  }

//================== OUTPUT ============================
  fmt.Println(ret)
}
