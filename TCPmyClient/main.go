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
    // Tạo một TCPAddr từ đối số cho chương trình
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    // Hàm DialTCP dùng để thiết lập một kết nối TCP trên client
    // trả về một *TCPConn và lỗi khi thực hiện
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)
    // method Write() của kiểu *TCPConn cho phép ghi vào (đưa vào) kết nối một slice byte
    for {
        var input string
        fmt.Println("Enter string: ")
        fmt.Scanln(&input)
        
        _, err = conn.Write([]byte(input))
        checkError(err)

        // func ReadAll của gói ioutil cho phép đọc một slice byte có trên kết nối *TCPConn
        buf := make([]byte, 1, 1)
        for {
            _, err := conn.Read(buf)
            checkError(err)
            fmt.Println(string(buf))
        }
    }
    
    os.Exit(0)
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}