package main

import (
    "net"
    "os"
    "fmt"
	"io"
)

func main()  {
    // địa chỉ service dưới kiểu chuỗi
    service := ":1201"

    // tạo một *TCPAddr
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    // là để lắng nghe trên net interfaces: port được đưa bởi đối số *TCPAddr
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    // liên tục bind việc chấp nhận kết nối từ client vào server ...
    for {
        // chấp nhận khi có kết nối từ client, trả về kiểu Conn
        conn, err := listener.Accept()
        if err != nil { // nếu không có
            continue
        }
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn)  {
    defer conn.Close()
    // tạo một mảng buf gồm 512 byte
    var buf [512]byte
    for { // vòng lập phiên
        // Method Read() của kiểu Conn đọc từ dữ liệu từ kết nối vào slice byte (được đưa vào từ đối số)
        // slice buf chỉ chứa tối đa 512 byte nên đọc tuần tự trong từng phiên 
        // các giá trị gán cho phần tử của slice cũng được gán cho mảng khởi tạo nên slice
        fmt.Println("waiting ...")
        n, err := conn.Read(buf[0:])
        fmt.Println("catched !")
         // n là số byte đọc vào thành công
        if err == io.EOF {
            fmt.Println("bye !")
            return
        }
        if err != nil {
            fmt.Println("error while handling client !")
            return
        }
        // in ra console dữ liệu đọc được
        fmt.Println(string(buf[0:n]))
        // đưa vào kết nối một slice byte
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}