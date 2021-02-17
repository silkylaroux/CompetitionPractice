package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  "strings"
)

func solveFromFile() int{
    //var nu int
    file, err := os.Open(os.Args[1])
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    //nu= fastGetInt(scanner.Bytes())

    defer file.Close()

    return solution(scanner.Text())
}

//========== Implement Algorithm =======================
func solution(num string) int{
  if strings.Contains(num, "Q")||strings.Contains(num, "H")||strings.Contains(num, "9") {
    return 1
  }else{
    return 0
  }
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
    //num = fastGetInt(scanner.Text())
    ret = solution(scanner.Text())
  }

//================== OUTPUT ============================
  if  ret == 1{
    fmt.Println("YES")
  }else{
    fmt.Println("NO")
  }
}
