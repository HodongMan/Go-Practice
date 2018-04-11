package main

import (
    "fmt"
    "os"
    "crypto/sha256"
    "crypto/sha512"
    "strconv"
)

func main() {

    if len(os.Args) == 2 {
	fmt.Printf("%x\n", MakeCrypto256(os.Args[1]))
    } else if len(os.Args) == 3 {
	version, err := strconv.Atoi(os.Args[2])

	if err == nil {

	    if version == 256 {
		fmt.Printf("%x\n", MakeCrypto256(os.Args[1]))
	    } else if version == 384 {
		fmt.Printf("%x\n", MakeCrypto384(os.Args[1]))
	    } else if version == 512 {
		fmt.Printf("%x\n", MakeCrypto512(os.Args[1]))
	    } else {
		fmt.Fprintf(os.Stderr, "crypto function is not exits")
	    }
	}

    } else {
	fmt.Fprintf(os.Stderr, "Argument must be <input> / <input> <crypto function>" )
    }
}

func MakeCrypto256(input string) [32]byte {    
    return sha256.Sum256([]byte(input))
}

func MakeCrypto384(input string) [48]byte {
    return sha512.Sum384([]byte(input))
}

func MakeCrypto512(input string) [64]byte {
    return sha512.Sum512([]byte(input))
}
