package main

import (
  "fmt"
  

                                                          /*
  "strings"
  "os"
  "bufio"
  "log"
  "strconv"
                                                          */
)


func main() {

//================== Variables used ====================
  var num, ret, p1, p2, p3 int

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

	  n, err2 := strconv.Atoi(scanner.Text())
    num = n
    for i := 0; i < num; i++ {
      scanner.Scan()
      p1,_ = strconv.Atoi(string(scanner.Text()[0]))
      p2,_ = strconv.Atoi(string(scanner.Text()[2]))
      p3,_ = strconv.Atoi(string(scanner.Text()[4]))
      if p1+p2+p3>=2{
        ret++
      }
    }
    if err2 != nil {
      log.Fatalf("failed opening file: %s", err2)
      
    }

//======================================================
    defer file.Close()
  }else{
                                                          */
//============== Assign vars from STDIN ================
    fmt.Scan(&num)
    for i := 0; i < num; i++ {
      fmt.Scan(&p1,&p2,&p3)
      if p1+p2+p3>=2{
        ret++
      }
    }
    
//======================================================
                                                          /*
  }
                                                          */
//========== Implement Algorithm =======================
  
//======================================================

//================== OUTPUT ============================
  fmt.Println(ret)
}
