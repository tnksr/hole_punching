/*
client
クライアント側
*/
package main

import (
    "log"
    "net"
    //"time"
    //"sync"
)

var otherAddr *net.UDPAddr 

func main() {
    // UDPのエンドポイントを返す
    //serverAddr, err := net.ResolveUDPAddr("udp", "13.59.148.159:60000") // EC2
    serverAddr, err := net.ResolveUDPAddr("udp", ":60000")  // local
    // エラー処理
    if err != nil {
       log.Fatal(err) 
    }

    // 接続開始
    serverConn, err := net.DialUDP("udp", nil, serverAddr)
    log.Println("dial up ...", serverAddr.String())
    if err != nil {
       log.Fatal(err) 
    }
    defer serverConn.Close()

    buffer := make([]byte, 128)
    
    for {
        // 自分のアドレスを送る
        myAddr := serverConn.LocalAddr()
        _, err := serverConn.Write([]byte(myAddr.String()))
        if err != nil {
            log.Fatal(err)
        }

        // 相手のアドレスを受け取る
        n, err := serverConn.Read(buffer)
        if err != nil {
            log.Fatal(err)
        }
        if n > 0 {
            receivedData, err := net.ResolveUDPAddr("udp", string(buffer[:n]))
            if err != nil {
                log.Fatal(err)
            }
            otherAddr = receivedData
            break
        }
    }

    // 受け取った相手のアドレスにリクエスト投げる
    for {
        userConn, err := net.DialUDP("udp", nil, otherAddr)
        if err != nil {
            log.Fatal(err)
        }
        defer userConn.Close()
    }
}
