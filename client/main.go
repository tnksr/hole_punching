package main

import (
    "log"
    "net"
    //"time"
)

func main() {
    // UDPのエンドポイントを返す
    //server_addr, err := net.ResolveUDPAddr("udp", "13.59.148.159:60000") // EC2
    server_addr, err := net.ResolveUDPAddr("udp", ":60000")  // local

    // エラー処理
    if err != nil {
       log.Fatal(err) 
    }
    // 接続開始
    conn, err := net.DialUDP("udp", nil, server_addr)
    log.Println("dial up ...")

    if err != nil {
       log.Fatal(err) 
    }
    defer conn.Close()

    buffer := make([]byte, 128)
    //other_addr := make(*net.UDPAddr, 8)
    var other_addr *net.UDPAddr
    
    for {
        // 自分のアドレスを渡す
        _, err := conn.Write([]byte(server_addr.String()))
        if err != nil {
            log.Fatal(err)
        }
        //time.Sleep(time.Second)
        // 相手のアドレスを受け取る
        n, err := conn.Read(buffer)
        if n > 0 {
            other_addr, err = net.ResolveUDPAddr("udp", string(buffer[:n]))
            log.Println(string(buffer[:n]), "buffer")
            log.Println(other_addr, "other")
            break
        }
    }

    // 受け取った相手のアドレスにリクエスト投げる
    for {
        break
    }
}
