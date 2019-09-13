// client 
package main

import (
    "log"
    "net"
    //"time"
    //"sync"
)

func main() {
    // UDPのエンドポイントを返す
    //serverAddr, err := net.ResolveUDPAddr("udp", "13.59.148.159:60000") // EC2
    serverAddr, err := net.ResolveUDPAddr("udp", ":60000")  // local
    // エラー処理
    if err != nil {
       log.Fatal(err) 
    }

    // 接続開始
    conn, err := net.DialUDP("udp", nil, serverAddr)
    log.Println("dial up ...", serverAddr.String())
    if err != nil {
       log.Fatal(err) 
    }

    buffer := make([]byte, 128)
    //otherAddr := make(*net.UDPAddr, 8)
    var otherAddr *net.UDPAddr
    
    for {
        // サーバに接続
        _, err := conn.Write([]byte(serverAddr.String()))
        if err != nil {
            log.Fatal(err)
        }
        //time.Sleep(time.Second)
        // 相手のアドレスを受け取る
        n, err := conn.Read(buffer)
        if n > 0 {
            otherAddr, err = net.ResolveUDPAddr("udp", string(buffer[:n]))
            log.Println(string(buffer[:n]), "buffer")
            log.Println(otherAddr, "other")
            break
        }
    }

    // 受け取った相手のアドレスにリクエスト投げる
    for {
        break
    }
}
