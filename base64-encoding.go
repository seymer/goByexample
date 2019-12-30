package main

import (
	bs64 "encoding/base64"
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnc := bs64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)


	sDec, _ := bs64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := bs64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := bs64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
