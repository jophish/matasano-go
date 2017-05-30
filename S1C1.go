package main

import (
	"./set1"
	"flag"
	"fmt"
	"os"
)

func main() {
	boolPtr := flag.Bool("b64", false, "Convert base64 string to hex string (default is hex->base64)")
	var svar string
	flag.StringVar(&svar, "str", "deadbeef", "string to convert")
	flag.Parse()
	if !*boolPtr {
		hex := set1.StrToHex(svar)
		base64, err := set1.HexToBase64(hex)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(set1.Base64ToStr(base64))
	} else {
		base64 := set1.StrToBase64(svar)
		hex, err := set1.Base64ToHex(base64)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(set1.HexToStr(hex))

	}

}
