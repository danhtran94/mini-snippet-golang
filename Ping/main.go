package main

import (
    "bytes"
    "fmt"
    "io"
    "net"
    "os"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "host")
        os.Exit(1)
    }

    // Tạo một IPAddr 
    addr, err := net.ResolveIPAddr("ip", os.Args[1])
    if err != nil {
        fmt.Println("Resolution error", err.Error())
        os.Exit(1)
    }

    // Tạo Raw Socket
    conn, err := net.DialIP("ip4:icmp", nil, addr)

    // tin chuyển
    var msg [512]byte 

    // format của ICMP
    msg[0] = 8
    msg[1] = 0
    msg[2] = 0
    msg[3] = 0
    msg[4] = 0
    msg[5] = 13
    msg[6] = 0
    msg[7] = 37
    len := 8

    // checkSum icmp
    check := checkSum(msg[0:len])
    msg[2] = byte(check >> 8)
    msg[3] = byte(check & 255)

    _, err = conn.Write(msg[0:len])
    _, err = conn.Read(msg[0:])

    fmt.Println("Got response")
    if msg[5] == 13 {
        fmt.Println("identifier matches")
    }
    if msg[7] == 37 {
        fmt.Println("sequence matches")
    }
    os.Exit(0)
}

// Hàm checkSum icmp
func checkSum(msg []byte) uint16  {
    sum := 0

    for n := 0; n < len(msg); n+=2 {
        sum += int(msg[n])*256 + int(msg[n+1])
        fmt.Println(sum)
    }

    sum = (sum >> 16) + (sum & 0xffff)
    fmt.Println(sum)
    sum += (sum >> 16)
    fmt.Println(sum)
    var answer = uint16(^sum)
    fmt.Println(answer)
    return answer
}

func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func readFully(conn net.Conn) ([]byte, error) {
    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
    }
    return result.Bytes(), nil
}