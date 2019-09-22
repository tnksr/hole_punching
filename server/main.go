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
    if _, ok := clientAddrMap[addr.String()]; !ok {
        clientAddrMap[addr.String()] = addr
    }
}

// 他のクライアントのアドレスを取得する
func getOtherAddr(clientAddr *net.UDPAddr) (*net.UDPAddr) {
    for addrString, otherAddr := range clientAddrMap {
        if addrString != clientAddr.String() {
            return otherAddr
        } 
    }
    return nil
}

func main() {
    // UDPのエンドポイントを返す
    serverAddr, err := net.ResolveUDPAddr("udp", ":60000")
    if err != nil {
       log.Fatal(err, "1: resolve udp address error.")
    }
    
    // 接続
    conn, err := net.ListenUDP("udp", serverAddr)
    if err != nil {
       log.Fatal(err, "2: listen udp error.")
    }
    defer conn.Close()
    
    buffer := make([]byte, 512)
    // 接続は2台までとする（仮）
    for len(clientAddrMap) < 2 {
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
    }

    for _, clientAddr := range clientAddrMap {
        // 接続クライアントに他のクライアントのアドレスを渡す
        otherAddr := getOtherAddr(clientAddr)
        log.Println(otherAddr)
        conn.WriteToUDP([]byte(otherAddr.String()), clientAddr)
    }
}


