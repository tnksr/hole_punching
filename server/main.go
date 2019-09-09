package main

import (
    "log"
    "net"
)


func main() {
    // UDPのエンドポイントを返す
    server_addr, err := net.ResolveUDPAddr("udp", ":60000")
    if err != nil {
       log.Fatal(err, "1")
    }
    
    // 接続確認
    conn, err := net.ListenUDP("udp", server_addr)
    if err != nil {
       log.Fatal(err)
    }
    defer conn.Close()

    log.Println(conn.RemoteAddr(), "remote")
    log.Println(conn.LocalAddr(), "local")

    buffer := make([]byte, 512)
    for {
        // 相手のアドレスを受け取る
        n, client_addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(client_addr.String(), "client")
        log.Println(string(buffer[:n]), "buffer") // 受け取り
        conn.WriteToUDP([]byte(server_addr.String()), client_addr)
    }
}
