/*
client
クライアント側
*/
package main

import (
    "flag"
    "log"
    "net"
    "bufio"
    //"time"
    //"sync"
)

var otherAddr *net.UDPAddr 

def ScanStdinByte() {
    stdin := bufio.NewScanner(os.Stdin)
    for stdin.Scan() {
        text := stdin.Text()
        return []byte(text)
    }
}

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

    // P2P通信したい相手のアドレス
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
        receivedFromServer := string(buffer[:n])
        otherAddr, err = net.ResolveUDPAddr("udp", receivedFromServer)
        if err != nil {
            log.Fatal(err)
        }
    }

    // 接続
    conn, err := net.DialUDP("udp", nil, otherAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    receivedData := make([]byte, 256)
    sendData := ScanStdinByte()
    for (receivedData == nil) {
        // 送信
        _, err := conn.Write(sendData)
        if err != nil {
            log.Fatal(err)
        }
        // 受信
        n, err := conn.Read(receivedData)
        if err != nil {
            log.Fatal(err)
        }
    }
}
