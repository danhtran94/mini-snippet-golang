package main

import (
    "net"
    "os"
    "fmt"
)

func main()  {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s network-type services\n", os.Args[0])
        os.Exit(1)
    }
    tranmissType := os.Args[1]
    service := os.Args[2]
    // Hàm LookupPort trả về port number kiểu int được dăng ký cho service dựa trên tranmission controll protocol
    // và service name được truyền vào
    port, err := net.LookupPort(tranmissType, service)
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(2)
    }
    fmt.Println("Service port", port)
    os.Exit(0)
}