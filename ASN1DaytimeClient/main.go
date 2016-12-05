package main

import (
    "bytes"
    "encoding/asn1"
    "fmt"
    "io"
    "net"
    "os"
    "time"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
    }
    service := os.Args[1]
    // call service dựa trên đối số gội chương trình
    conn, err := net.Dial("tcp", service)
    checkError(err)
    // hàm đọc kết nối
    result, err := readFully(conn)
    checkError(err)
    // tạo một biến kiểu Time của package time
    var newTime time.Time
    // Unmarshal dữ liệu fill vào biến newTime
    // Hàm Unmarshal của gói asn1 phân tích cú pháp dữ liệu ASN1 đã mã hóa và sử dụng gói reflect để lắp đầy giá trị được trỏ tới
    _, err1 := asn1.Unmarshal(result, &newTime)
    checkError(err1)
    // in ra giá trị tái cấu trúc thành công
    fmt.Println("After marshal/unmarshal:", newTime.String())
    os.Exit(0)
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(0)
    }
}

func readFully(conn net.Conn) ([]byte, error) {
    defer conn.Close()
    // tạo một buffer giữ kết quả không giới hạn size nil
    result := bytes.NewBuffer(nil)
    // tạo một buffer chứa 512 byte dùng để chia đoạn dữ liệu trên kết nối
    var buf [512]byte
    for { // lặp phiên liên tục
        // đọc mỗi phiên tối đa 512 byte
        n, err := conn.Read(buf[0:])
        // ghi vào buffer kết quả dữ liệu vừa đọc
        result.Write(buf[0:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
    }
    // trả vể buffer kết quả dạng slice byte, nil là không lỗi xảy ra
    return result.Bytes(), nil
}