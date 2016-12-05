package main

import (
	"fmt"
	"crypto/md5"
	// "crypto/hmac"
)

func main() {
	// hash := hmac.New(md5.New, []byte("my key"))
	hash := md5.New()
	bytes := []byte("hello\n")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	fmt.Println(hashValue)
	hashSize := hash.Size()
	fmt.Println(hashSize)
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n]) << 24 +
		uint32(hashValue[n+1]) << 16 +
		uint32(hashValue[n+2]) << 8 +
		uint32(hashValue[n+3])
		fmt.Printf("%x.", val)
	}
	fmt.Println()
}