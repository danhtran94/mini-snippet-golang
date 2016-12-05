package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
	"io"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host:port")
		os.Exit(1)
	}
	
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	
	shorts := readShorts(conn)
	ints := utf16.Decode(shorts)
	str := string(ints)
	
	fmt.Println(str)
	os.Exit(0)
	
}

func readShorts(conn net.Conn) []uint16 {
	var buf [512]byte
	n, err := conn.Read(buf[0:2])
	for true {
		m, err := conn.Read(buf[n:])
		if err != nil {
			if err == io.EOF {
				fmt.Println("Kết thúc nhận thành công !")
			}
			break
		}
		n += m
	}
	
	checkError(err)
	var shorts []uint16
	shorts = make([]uint16, n/2)
	if buf[0] == 0xff && buf[1] == 0xfe {
		fmt.Println("Encoding kiểu Big Endien")
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i]) << 8 + uint16(buf[i+1])
		}
	} else if buf[1] == 0xff && buf[0] == 0xfe {
		fmt.Println("Encoding kiểu Little Endien")
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i+1])<<8 + uint16(buf[i])
		}
	} else {
		fmt.Println("Unknown order")
	}
	return shorts
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

