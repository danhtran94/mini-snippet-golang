package main

import (
    "net"
    "os"
    "fmt"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
        os.Exit(1)
    }
    name := os.Args[1]
    // Hàm LookupCNAME trả về canonical hostname (hostname chính) cùa hostname
    cname, err := net.LookupCNAME(name)
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(2)
    }
    // Hàm LookupHost trả về tất cả ip address của hostname dưới dạng string và trả về lỗi 
    addrs, err := net.LookupHost(name)
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(2)
    }
    fmt.Println(cname)
    for _, s := range addrs {
        fmt.Println(s)
    }
    os.Exit(0)
}