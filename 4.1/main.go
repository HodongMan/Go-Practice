package main

import (
    "fmt"
    "crypto/sha256"
)

func DiffBitCount(s1, s2 [32]byte) int {

    count := 0
    fmt.Printf("%x\n", s1)
    fmt.Printf("%x\n", s2)
    for index, item := range s1 {
	if item != s2[index] {
	    count += 1;
	}
    }

    return count
}

func main() {
    s1 := sha256.Sum256([]byte("abc"))
    s2 := sha256.Sum256([]byte("def"))

    result := DiffBitCount(s1, s2)

    fmt.Println(result)
}

