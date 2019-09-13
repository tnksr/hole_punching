/*
server
    サーバ側
*/
package main

import (
    "log"
    "net"
    "sync"
)

// クライアントのアドレスを格納しておくハッシュ
var (
	clientAddrMap = make(map[string]*net.UDPAddr, 8)
	mutex      sync.Mutex
)

// クライアントのアドレスを登録する
func registClientAddr(addr *net.UDPAddr) {
    mutex.Lock()
    defer mutex.Unlock()
    clientAddrMap[addr.String()] = addr
}

func main() {
    // UDPのエンドポイントを返す
    serverAddr, err := net.ResolveUDPAddr("udp", ":60000")
    if err != nil {
       log.Fatal(err, "1: resolve udp address error.")
    }
    
    // 接続確認
    conn, err := net.ListenUDP("udp", serverAddr)
    if err != nil {
       log.Fatal(err, "2: listen udp error.")
    }
    defer conn.Close()
    
    buffer := make([]byte, 512)
    for {
        // クライアントのアドレスを受け取る
        _, clientAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            log.Fatal(err, "3: connect error")
        }
        // 接続クライアントのアドレス
        // log.Println(conn.RemoteAddr(), "remote")
        // サーバのアドレス 
        // log.Println(conn.LocalAddr(), "local")

        // 接続クライアントを登録
        log.Println(clientAddrMap)
        registClientAddr(clientAddr)
        log.Println(clientAddrMap)

        /* TODO
        クライアントにアドレスを渡す
        今はサーバのアドレス渡してるので
        clientAddrMapから取り出して渡す
        */
        conn.WriteToUDP([]byte(serverAddr.String()), clientAddr)
    }
}


