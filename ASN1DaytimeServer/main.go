package main

import (
    "encoding/asn1"
    "fmt"
    "net"
    "os"
    "time"
)

func main()  {
    // địa chỉ server dạng chuỗi
    service := ":1200"
    // tạo địa chỉ kiểu TCPAddr 
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)
    // lắng nghe trên interface:port
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    fmt.Println("Listening ...")

    for {
        // chấp nhận kết nối tới trả về một interface Conn
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        // lấy thời gian hiện tại
        dayTime := time.Now()
        // Bỏ qua kiểm tra trường hợp mạng bị lỗi tại dòng này
        // Hàm Marshal của gói asn1 trả về giá trị được mã hóa ASN1 dạng slice byte
        mdata, _ := asn1.Marshal(dayTime)
        // ghi dữ liệu ra kết nối
        conn.Write(mdata)
        // đóng kết nối
        conn.Close()
    }
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}