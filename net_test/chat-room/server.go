package main

import (
    "bufio"
    "fmt"
    "net"
    "log"
)

var ConnMap map[string]*net.TCPConn

func main() {
    var tcpAddr *net.TCPAddr
    ConnMap = make(map[string]*net.TCPConn)
    tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
    checkError(err)

    
    tcpListener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    
    defer tcpListener.Close()
    
    for {
        tcpConn, err := tcpListener.AcceptTCP()
        if err != nil {
            continue
        }
        
        fmt.Println("A Clinet connected: " + tcpConn.RemoteAddr().String())
        ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
        go tcpPipe(tcpConn)
    }
}

func tcpPipe(conn *net.TCPConn) {
    ipStr := conn.RemoteAddr().String()
    defer func() {
        fmt.Println("disconnected:" + ipStr)
        conn.Close()
    }()
    reader := bufio.NewReader(conn)
    
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            return 
        }
        fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
        boradcastMessage(conn.RemoteAddr().String() + ":" + string(message))
    }
}

//直接遍历广播消息
func boradcastMessage(message string) {
    b := []byte(message)
    for _, conn := range ConnMap {
        conn.Write(b)
    }
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
