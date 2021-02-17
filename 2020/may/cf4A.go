package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
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
    var nu int
    file, err := os.Open(os.Args[1])
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    nu= fastGetInt(scanner.Bytes())

    defer file.Close()

    return solution(nu)
}

//========== Implement Algorithm =======================
func solution(num int) int{
  return num%2
}

func main() {

//================== Variables used ====================
  var num, ret int
//======================================================

  if len(os.Args) >=2{
//============== Get return value from File ============
    ret = solveFromFile()


  }else{
//============== Get return value from STDIN ===========
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    num = fastGetInt(scanner.Bytes())
    ret = solution(num)
  }

//================== OUTPUT ============================
  if ret==0 && num > 2 {
    fmt.Println("Yes")
  }else{
    fmt.Println("No")
  }
  
}
