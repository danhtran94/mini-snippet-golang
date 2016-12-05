package main

import (
    "net"
    "os"
    "fmt"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)

    conn, err := net.DialUDP("udp", nil, udpAddr)
    checkError(err)
    err = conn.SetReadBuffer(512)
    checkError(err)

    _, err = conn.Write([]byte("Anything"))
    checkError(err)

    var buf [512]byte
    n, err := conn.Read(buf[0:])
    checkError(err)

    fmt.Print(string(buf[0:n]))
    os.Exit(0)
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error", err.Error())
        os.Exit(1)
    }
}