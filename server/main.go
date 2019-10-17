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
// マップのキーを取り出す
func keys(m map[string]*net.UDPAddr) []string {
    ks := []string{}
    for k, _ := range m {
        ks = append(ks, k)
    }
    return ks
}

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
    clientNum := 2
    for {
        for len(clientAddrMap) < clientNum {
            // クライアントのアドレスを受け取る
            _, clientAddr, err := conn.ReadFromUDP(buffer)
            if err != nil {
                log.Fatal(err, "3: connect error")
            }
            // 接続クライアントを登録
            registClientAddr(clientAddr)
            // 登録したら表示
            log.Println(keys(clientAddrMap))
        }

        // 接続クライアントに他のクライアントのアドレスを渡す
        for _, clientAddr := range clientAddrMap {
            otherAddr := getOtherAddr(clientAddr)
            conn.WriteToUDP([]byte(otherAddr.String()), clientAddr)
        }
        // mapを空にする
        clientAddrMap = make(map[string]*net.UDPAddr)
        // 削除したら表示
        log.Println(keys(clientAddrMap))
    }
}
