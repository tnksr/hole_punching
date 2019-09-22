/*
client
クライアント側
*/
package main

import (
    "flag"
    "log"
    "net"
    //"time"
    //"sync"
)

var otherAddr *net.UDPAddr 

func main() {
    // サーバのアドレス
    var serverAddrString string
    portString := ":60000"
    // サーバのアドレスをコマンドライン引数から受け取る
    flag.Parse()
    parseArgs := flag.Args()
    if len(parseArgs) < 1 {
        serverAddrString = ""
    } else {
        serverAddrString = parseArgs[0]
    }
    serverAddrString += portString

    // サーバを探す
    serverAddr, err := net.ResolveUDPAddr("udp", serverAddrString)
    if err != nil {
       log.Fatal(err) 
    }

    // 接続開始
    serverConn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
       log.Fatal(err) 
    }
    defer serverConn.Close()
    log.Println("Connect: ", serverAddr.String())

    for (otherAddr == nil) {
        // 自分のアドレスを送る
        myAddr := serverConn.LocalAddr()
        _, err := serverConn.Write([]byte(myAddr.String()))
        if err != nil {
            log.Fatal(err)
        }

        // 相手のアドレスを受け取る
        buffer := make([]byte, 128)
        n, err := serverConn.Read(buffer)
        if err != nil {
            log.Fatal(err)
        }

        // 受け取ったアドレスを探す
        receivedData := string(buffer[:n])
        otherAddr, err = net.ResolveUDPAddr("udp", receivedData)
        if err != nil {
            log.Fatal(err)
        }
    }
}
