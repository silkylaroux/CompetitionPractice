package main

import (
	"fmt"
	"strings"
    "bufio"
    "os"
)

func main() {
    st:= ""
	
	in := bufio.NewReader(os.Stdin)

    fmt.Fscan(in, &st)
    stnew :=""
    
    
    //st = strings.Replace(st,"m","",-1)
    //st = strings.Replace(st,string(st[0]),"",-1)
    for len(st)>0{
       stnew += string(st[0])
       st = strings.Replace(st,string(st[0]),"",-1)
    }
    //stnew = st
    if(len(stnew)%2)==0{
        fmt.Println("CHAT WITH HER!")
    }else{
        fmt.Println("IGNORE HIM!")
    }
    
}
