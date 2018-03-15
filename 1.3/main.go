
package main

import(
    "time"
    "fmt"
    "os"
    "strings"
)

func main(){

    var s, step string

    start := time.Now()
    for i:=0; i < len(os.Args); i++ {

	s += step + os.Args[i]
	step = " "
    }

    fmt.Println(s)
    fmt.Println("완료 시간 : ", time.Since(start).Seconds())
    
    start2 := time.Now()

    fmt.Println(strings.Join(os.Args, " "))
    fmt.Println("완료 시간 : ", time.Since(start2).Seconds())
}
