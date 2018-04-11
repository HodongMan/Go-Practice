package main

import (
    "fmt"
    "os"
    "bytes"
)

func main() {
    for i := 1; i < len(os.Args); i++ {
	fmt.Printf("%s\n", comma(os.Args[i]))
    }
}

func comma(s string) string {

    var buf bytes.Buffer

    for i := 1; i < len(s); i++ {
	
	if (i % 3 == 0) {
	    buf.WriteString(s[i-3:i])
	    buf.WriteString(", ")
	}
    }
    return buf.String()
}
