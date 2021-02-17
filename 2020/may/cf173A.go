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
  var num, ret int
  var line string
  
//======================================================
                                                          /*
    hasFile:=false
//============== Assign vars from file =================
  scanner := bufio.NewScanner(os.Stdin)

  if strings.Compare(os.Args[1], "") != 0{
    hasFile = true
    file, err := os.Open(os.Args[1])
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    scanner = bufio.NewScanner(file)
    scanner.Scan()

	n, err2 := strconv.Atoi(scanner.Text())
    num = n
    if err2 != nil {
      log.Fatalf("failed opening file: %s", err2)
      
    }

//======================================================
    defer file.Close()
  }else{
                                                          */
//============== Assign vars from STDIN ================
    fmt.Scan(&num)
    
    //fmt.Scan(&line)
//======================================================
                                                          /*
  }
                                                          */
//========== Implement Algorithm =======================
  for x := 0;x<num; x++{
                                                          /*
    if hasFile{
      scanner.Scan()
      if string(scanner.Text()[0]) == "-"||string(scanner.Text()[1]) == "-"{
        ret--
      }else{
        ret++
      }
    }else{                                              */
      fmt.Scan(&line)
      
      if string(line[0]) == "-"||string(line[1]) == "-"{
        ret--
      }else{
        ret++
      }
    }                                                   /*
  }                                                     */
//======================================================

//================== OUTPUT ============================
  fmt.Println(ret)
}
