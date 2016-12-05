package main

import (
    "net"
    "os"
    "fmt"
)
func main()  {
    // len() trả về độ dài slices
    // os.Args là slice chứa các đối số khi gọi chương trình
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
        // os.Exit thoát chương trình kèm theo tín hiệu thoát 1 là có lỗi xảy ra
        os.Exit(1)
    }
    name := os.Args[1]
    // net.ParseIP() nhận một địa chỉ ip dạng string chuyển về kiểu net.IP
    addr := net.ParseIP(name)
    if addr == nil {
        fmt.Println("Invalid address")
    } else {
        fmt.Println("The address is", addr.String()) // method .String của type net.IP trà về IP đó dạng chuỗi
    }
    os.Exit(0)
}