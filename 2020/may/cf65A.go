package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  //"strconv"
)
//similiar functions on internet. 
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

func fastGetStr(b byte) string{
  t := ""
	if b >= 'A'&& b<='Z' {
		t = string('Z'-('Z'-b))
  }
  if b >= 'a'&& b<='z' {
		t = string('z'-('z'-b))
	}
	
	return t
}

func solveFromFile() {
    var nu int
    file, err := os.Open(os.Args[1])
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    nu= fastGetInt(scanner.Bytes())
    for x := 0; x < nu; x++ {
      scanner.Scan()
      solution(scanner.Bytes())
    }
    defer file.Close()

}

//========== Implement Algorithm =======================
func solution(num []byte){
  if len(num)>10 {
    //fmt.Println(num)
    first:= num[0]
    mid:= (len(num)-2)
    last := (num[len(num)-1])
    fmt.Printf(fastGetStr(first))
    fmt.Printf("%d",mid)
    fmt.Printf(fastGetStr(last))
    fmt.Println()

    //fmt.Println(first+mid+last)
    //fmt.Sprintf("%s%s%s",num[0],len(num)-2,num[len(num)-1])
   // fmt.Println(string(num[0])+string(len(num)-2)+string(num[len(num)-1]))
  }else{
    fmt.Println(string(num))
  }
}

func main() {

//================== Variables used ====================
  var num int
//======================================================

  if len(os.Args) >=2{
//============== Get return value from File ============
    solveFromFile()


  }else{
//============== Get return value from STDIN ===========
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    num = fastGetInt(scanner.Bytes())
    for x := 0; x < num; x++ {
      scanner.Scan()
      solution(scanner.Bytes())
    }
  }

//================== OUTPUT ============================
  //fmt.Println(ret)
}
