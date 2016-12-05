package main

import (
    "net"
    "os"
    "fmt"
    "io/ioutil"
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
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    // func ReadAll của gói ioutil cho phép đọc một slice byte có trên kết nối *TCPConn
    result, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))

    os.Exit(0)
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}