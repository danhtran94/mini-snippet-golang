package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main()  {
    service := ":1200"
    // Tạo một TCPAddr trường hợp này là localhost port 1200
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    // Hàm ListenTCP() của gói net lắng nghe trên network interface(s):port, trả về một *TCPListener và lỗi 
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    // liên tục bind vào việc chấp nhận kết nối cho client đến server
    for {
        // Method Accept() của kiểu *TCPListener chấp nhập kết nối đến các port đang nghe ngóng, trả về một kiểu Conn và lỗi
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        // Hàm Now() của package time trả về thời gian hiện tại của máy tính kiểu Time
        // Method String() của kiểu Time trả về thời gian dưới dạng chuỗi
        daytime := time.Now().String()
        // Method Write() của kiểu Conn cho phép đưa vào (ghi vào) kết nối một slice byte, trả về số nguyên n và lỗi
        conn.Write([]byte(daytime+"\n"))
        // Method Close() của kiểu Conn dùng để kết thúc kết nối
        conn.Close()
    }
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}