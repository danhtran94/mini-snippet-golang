package main

import (
	"fmt"

)

var unicodeToISOMap = map[rune]uint8{
	0x12e: 0xc7,
	0x10c: 0xc8,
	0x118: 0xca,
}

func main()  {
	uniStr := string([]rune{0x12e, 0x10c, 0x118})
	isoArr := unicodeStrToISO(uniStr)
	uniStr = isoBytesToUnicode(isoArr)
	fmt.Println(uniStr)
}

func unicodeStrToISO(str string) []byte {
	codePoints := []rune(str)
	bytes := make([]byte, len(codePoints))
	for n, v := range codePoints {
		iso, ok := unicodeToISOMap[v]
		if !ok {
			iso = uint8(v)
		}
		bytes[n] = iso
	}
	return bytes
}

var isoToUnicodeMap = map[uint8] rune {
	0xc7: 0x12e,
	0xc8: 0x10c,
	0xca: 0x118,
	// and more
}

func isoBytesToUnicode(bytes []byte) string {
	codePoints := make([]rune, len(bytes))
	for n, v := range(bytes) {
		unicode, ok := isoToUnicodeMap[v]
		if !ok {
			unicode = rune(v)
		}
		codePoints[n] = unicode
	}
	return string(codePoints)
}