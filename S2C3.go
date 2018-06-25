package main

import (
	"./set1"
	"fmt"
	"os"
)

// We want to detect what AES mode set1.AES_EncryptionOracle uses. We can craft our plaintext
// to make this job easier.
func main() {
	l := 100
	buf := make([]byte, l)
	for i := 0; i < l; i++ {
		buf[i] = 0
	}
	cipher := set1.AES_EncryptionOracle(buf)
	//fmt.Println(cipher)
	for i := 0; i < 16; i++ {
		if cipher[16+i] != cipher[32+i] {
			fmt.Println("CBC mode detected!")
			os.Exit(0)
		}

	}
	fmt.Println("ECB mode detected!")

}
