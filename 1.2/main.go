
package main

import(
    "fmt"
    "os"
)

func main(){

    for i:=0; i < len(os.Args); i++ {

	fmt.Println(i + 1, " : ", os.Args[i])
    }

}
