package main

import (
	"./set1"
	"fmt"
)

func main() {
	str1 := "this is a test"
	str2 := "wokka wokka!!!"

	buf1 := set1.StrToASCII(str1)
	buf2 := set1.StrToASCII(str2)
	ans, _ := set1.HammingDistance(buf1, buf2)
	fmt.Println(ans)

	str3 := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	buf3 := set1.StrToHex(str3)
	guess := set1.GuessKeySize(buf3)
	fmt.Println(guess)
}
