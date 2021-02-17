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
    nu= solution(scanner.Bytes())
    //fastGetInt(scanner.Bytes())

    defer file.Close()

    return nu
}

//========== Implement Algorithm =======================
func solution(num []byte) int{
  count:=0
  prev := num[0]
  for i := 1; i < len(num); i++ {
    if prev == num[i]{
      count++
    }else{
      count=0
    }
    if count==6 {
      return 0
    }
    prev = num[i]
  }
  return 1
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
    num = solution(scanner.Bytes())
    //fastGetInt(scanner.Bytes())
    ret = num//solution(num)
  }

//================== OUTPUT ============================
  if ret==0 {
    fmt.Println("YES")
  }else{
    fmt.Println("NO")
  }
  //fmt.Println(ret)
}
