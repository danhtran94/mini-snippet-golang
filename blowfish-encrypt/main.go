package main

import (
	"bytes"
	"golang.org/x/crypto/blowfish"
	"fmt"
	
)

func main() {
	key := []byte("1721119")
	fmt.Println("Mã hóa và giải mã với key:", string(key))
	cipher, err := blowfish.NewCipher(key)
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	src := []byte("Danh yêu Dung")
	length := len(src)
	fmt.Println("Độ dài src trước khi bù:", length)
	
	if residual := length % 8; residual != 0 {
		pad := 8 - residual
		for count := 0; count < pad; count++ {
			src = append(src, byte(0x20))
		}
	}
	
	length = len(src)
	fmt.Println("Độ dài src sau khi bù:", length)
	var enc [8]byte
	encResult := bytes.NewBuffer(nil)
	
	times := length / 8
	fmt.Println("Số lần để mã hóa hết src:", times)
	n := 1
	for i := 0; n <= times; i += 8 {
		fmt.Println("Mã hóa (8 bytes) lần thứ", n, ", từ byte thứ", i, "->", i + 7)
		cipher.Encrypt(enc[0:], src[i:i + 8])
		encResult.Write(enc[0:])
		fmt.Println("Xong !")
		n++
	}
	
	fmt.Print("Sau khi src được mã hóa: ")
	for _, byte := range enc {
		fmt.Printf("%x", byte)
	}
	fmt.Println()
	
	result := bytes.NewBuffer(nil)
	var decrypt [8]byte
	m := 1
	er := encResult.Bytes()
	for i := 0; m <= times; i += 8{
		cipher.Decrypt(decrypt[0:], er[i:i + 8])
		result.Write(decrypt[0:])
		m++
	}
	fmt.Println("Độ dài sau khi giải mã:", len(result.Bytes()))
	fmt.Println(string(result.Bytes()))
}
