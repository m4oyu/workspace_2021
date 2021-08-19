# chap 2 インストールと起動

## 2.1 ソースコードからのインストール

## 2.2 パッケージからのインストール

参照：http://nginx.org/en/linux_packages.html#Ubuntu
- 公式の署名キー（GPGキー）のインポートがうまくいかなかったが、無視してインストールした。

```
~$ nginx -V
nginx version: nginx/1.18.0 (Ubuntu)
built with OpenSSL 1.1.1f  31 Mar 2020
TLS SNI support enabled
configure arguments: --with-cc-opt='-g -O2 -fdebug-prefix-map=/build/nginx-KTLRnK/nginx-1.18.0=. ...
```

### GPGキー
`通常，Linux環境では自分が意図した作成者が作ったソフトであるか、
使うインストーラーが本物である事を検証（身元が正しいかどうか）してからインストールします．
このとき，ソフトウェアの検証に使用される仕組みの一つがgpgです．`\
`gpgはGNU Privacy Guard (GnuPG, GPG)の略で暗号化ソフトウェアです.`\
`暗号化だけでなく署名や認証といったオンライン上の機密や信用を管理するツールとしても利用されます．特徴として公開鍵暗号方式を使用しますが，その鍵の管理のために，認証局を設置しません．各利用者の責任で鍵を管理し，取得した公開鍵をチェックします．`
[引用](https://qiita.com/y518gaku/items/435838097c700bbe6d1b)


## 2.3 nginxの起動、終了、基本的な操作

### nginxの起動
- rootユーザ権限でないと失敗する
```
sudo nginx
```
- nginx -t で nginx.confファイルの構文チェックを行う
```
$ sudo nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
```
- masterプロセスとworkerプロセスを起動する
- 80番ポートをバインド

### nginxの終了、設定の彩度読み込み
- nginxコマンドによる制御
  - -s オプションによってマスタプロセスへシグナルを飛ばす
  - stop : リクエストの処理を待たずに終了
  - quit : 現在の処理後、終了
  - reload : 新しい設定ファイルの読み込み、プロセスの引継ぎはよしなに
```
sudo nginx -s stop
```
- kill コマンドによるシグナルの送信
```
kill -s QUIT `cat /var/run/nginx.pid`
```

### nginxをシステムサービスとして実行
- システムサービス : 常駐するプログラム
- サービス関連のコマンド
```
sudo service nginx start
sudo service nginx stop
etc...
```
- systemctlによるコマンド
```
sudo systemctl enable nginx.service
etc...
```

## appendix
### tee コマンド
「tee」は標準入力から受け取った内容を、標準出力とファイルに書き出すコマンドです。ファイルへの保存と標準出力への出力を同時に行ったり、複数のファイルに出力したりすることができます。
- 標準入力のリダイレクトとパイプを同時に行う
```
コマンド1 | tee ファイル | コマンド2
```
- 複数のファイルへ出力する
```
コマンド | tee ファイル1 ファイル2 ファイル3……
```

### ss コマンド
「ss」コマンドは、ネットワーク通信で利用する「ソケット」についての情報などを出力するコマンドです。\
従来はnetstatコマンドが使用されていましたが、現在はssコマンドへの移行が進んでいます。