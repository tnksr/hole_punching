# Hole Punching
- golangを触ってみる
- UDP通信のおべんきょう

# 目標
- 2019/09/06 社内LT
- この日までに通信したい

# ToDo
- client側のIPアドレスの入力
    - [x] インスタンス立ち上げるごとにアドレス変わるのめんどう（とりあえずどうしようもなさそう／質問済み）
- 通信メッセージをお互いにやりとり
    - [x] 今はbufferで0ばっか出る
    - [ ] 通信日時とアドレスとか出ればいいかな
    - [ ] アドレスのやりとり


# Setting

## EC2の立ち上げ

    1. キーペアの作成（ダウンロードしておく）
    2. セキュリティグループに
        - タイプ：カスタムUDPルール
        - プロトコル：UDP
        - ポート範囲：ソースコードに合わせる（今回は60000）
        - ソース：立ち上げたインスタンスのIP
       を設定
    3. runningにする
    4. ローカルで
        
        `$ ssh ec2-user@ec2-X-XXX-XX-XX.us-east-2.compute.amazonaws.com -i /path/to/downloaded/pem`

## Goのインストール
    1. 公式ページからダウンロード
        
        `$ curl -O https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz`

    2. 解凍

        `$ sudo tar -C /usr/local -xzf go1.12.9.linux-amd64.tar.gz`

    3. bash_profileの設定
        
        ``
        export GOPATH=$HOME/go
        export PATH=$PATH:/usr/local/go/bin
        ``

    4. 確認
        
        ``
        $ go version
        ``
