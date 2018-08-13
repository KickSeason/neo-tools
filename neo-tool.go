package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/CityOfZion/neo-go/pkg/crypto"
)

func toHash160(address string) string {
	u, err := crypto.Uint160DecodeAddress(address)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	rs := hex.EncodeToString(u.BytesReverse())
	return rs
}
func toDecimal(hex string) string {
	if len(hex)%2 != 0 {
		return ""
	}
	if len(hex) == 2 {
		n, err := strconv.ParseInt(hex, 16, 64)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		r := strconv.Itoa(int(n))
		return r
	}
	hexBytes := []byte(hex)
	destBytes := []byte{}
	for i := 0; i < len(hexBytes); i += 2 {
		destBytes = append(destBytes, hexBytes[len(hexBytes)-i-2])
		destBytes = append(destBytes, hexBytes[len(hexBytes)-i-1])
	}
	n, err := strconv.ParseInt(string(destBytes), 16, 64)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	r := strconv.Itoa(int(n))
	return r
}
func main() {
	if len(os.Args) < 3 {
		return
	}
	if os.Args[1] == "-a" {
		fmt.Printf("addr: %s \nhash160: %s\n", os.Args[2], toHash160(os.Args[2]))
	} else if os.Args[1] == "-h" {
		fmt.Printf("hex: %s \nhash160: %s\n", os.Args[2], toDecimal(os.Args[2]))
	}
}
