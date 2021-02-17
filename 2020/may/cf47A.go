package main

import (
  "fmt"
  
                                                          /*
  "strconv"
  "os"
  "bufio"
  "log"
  "strings"
  
                                                          */
)


func main() {

//================== Variables used ====================
  var num, tempNum1,tempNum2 int
  //var nums string

//======================================================
                                                          /*
//============== Assign vars from file =================
  scanner := bufio.NewScanner(os.Stdin)

  if strings.Compare(os.Args[1], "") != 0{

    file, err := os.Open(os.Args[1])
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    scanner = bufio.NewScanner(file)
    scanner.Scan()
    n, err2 := strconv.Atoi(string(scanner.Text()[0]))
  
    num = n
    n2, err3 := strconv.Atoi(string(scanner.Text()[2]))
  
    num = num * n2
    if err2 != nil || err3 != nil {
      log.Fatalf("failed opening file: %s", err2)
      
    }

//======================================================
    defer file.Close()
  }else{
                                                          */
//============== Assign vars from STDIN ================
    fmt.Scan(&tempNum1)
    fmt.Scan(&tempNum2)
    
//======================================================
                                                          /*
  }
                                                          */
//========== Implement Algorithm =======================


  num = tempNum1*tempNum2

//======================================================

//================== OUTPUT ============================
  fmt.Println(int(num/2))
}
