package main

import (
    "fmt"
    "net"
    "os"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
        os.Exit(1)
    }
    dotAddr := os.Args[1]
    addr := net.ParseIP(dotAddr)
    if addr == nil {
        fmt.Println("Invalid address")
        os.Exit(1)
    }
    // Phương thức DefaultMask() của kiểu IP trả về mask mặc định có kiểu IPMark
    mask := addr.DefaultMask()
    // Phương thức Mask(mask IPMark) của kiểu IP trả về IP addr của network có kiểu IP 
    network := addr.Mask(mask)
    // Phương thức Size() của kiểu IPMask trả về số bit cho phần mạng và tổng số bit cùa IPMark chúng có kiểu int
    ones, bits := mask.Size()
    // Phương thức String() của kiểu IPMark trả về chuỗi thể hiện một số thập lục phân
    fmt.Println("Address is", addr.String(), 
    "\nDefault mask length is", bits, 
    "\nLeading ones count is", ones, 
    "\nMask is (hex)", mask.String(),
    "\nNetwork is", network.String())
    os.Exit(0)
}