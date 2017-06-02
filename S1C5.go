package main

import (
	"./set1"
	"fmt"
)

func main() {
	str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	strBuf := set1.StrToASCII(str)
	key := "ICE"
	keyBuf := set1.StrToASCII(key)
	encrypted := set1.RepeatingXOR(strBuf, keyBuf)
	fmt.Println(set1.HexToStr(encrypted))
	decrypted := set1.RepeatingXOR(encrypted, keyBuf)
	fmt.Println(string(decrypted))
}
