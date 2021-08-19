# chap 3 基本設定

# 3.1 設定ファイルの構成

## 設定ファイルのフォーマット
## ディレクティブ
## 変数
- $exampleで変数を宣言できる
## 設定のインクルード
- includeディレクティブで設定ファイルを読み込める
  - 複数サーバを動かす際に便利
  - 共通設定を一つのファイルにまとめられる
  - nginx.confファイルからの相対パスもしくは絶対パス
  - includeディレクティブの場所にそのまま読み込み内容が書き込まれる

# 3.2 HTTPサーバに関する設定
- nginxは複数のモジュールからなる
- HTTPサーバに必要な機能はngx_http_core_moduleに実装されている

## HTTPコンテキストの定義
- HTTPサーバの動作に関する設定はほとんどHTTPコンテキスト中で定義
- httpディレクティブによって定義される
```
http {
    tcp_nopush on;
    ...
}
```

## バーチャルサーバの定義
- nginxでは使用するIPアドレス、ポート、ホスト名ごとにHTTPサーバを動作させることができる（バーチャルサーバ）
- serverディレクティブで定義する
```
http {
    server {
        listen 80;
        server_name www.example.com
        ...
    }
    server {
        listen 80;
        server_name www.ex2.com
        ...
    }
}
```
- listenディレクティブ : 使用するアドレス、ポートの指定
  - unixドメインソケット（linux内でのプロセス間通信に使用される）も指定できる
- server_nameディレクティブ : バーチャルサーバで使用するホスト名を指定する

### 複数のバーチャルサーバの優先順位
1. listenディレクティブのアドレスとポートに一致するバーチャルサーバを検索する
2. リクエストのHostヘッダがserver_nameディレクティブで指定したホストに一致下バーチャルサーバに振り分ける
3. どのサーバにも一致しない場合、デフォルトサーバに振り分ける

- デフォルトサーバ
```
server {
    listen 80 default_server;
}
```

## 公開するディレクトリを設定
- rootディレクティブ : 公開するディレクトリを明確に指定する
- /var/www/htmlというディレクトリを公開する
```
root /var/www/html;
```

## MIMEタイプの指定
- typesディレクティブ : MIMEタイプとそれに対応する拡張子の指定を行う
  - MIMEタイプ : Content-Typeヘッダフィールドでファイルの指定を行うもの
  - nginxには標準的なMIMEタイプを指定したmime.typesファイルが添付されている。  
  それをインクルードするだけで基本は大丈夫
```example
include /etc/nginx/mime.type
```
```
$ cat mime.type

types {
    text/html                             html htm shtml;
    text/css                              css;
    ...
}
```

### default_typeディレクティブ
| defaut_type | defaultのMIMEタイプを指定する |
| ----------- | ----------------------------- |
| 構文        | default_type　MIMEタイプ;     |
| default     | text/plain                    |
| context     | http, server, location        |

- types {} : ｛｝内を空にすることですべてのマッピングを無効化する
- default_type application/octet-stream : すべてのファイルがブラウザでダウンロードするように振る舞う
```
types {}
default_type application/octet-stream
```

## アクセスログの出力
- アクセスログに関するディレクティブはngx_http_log_moduleに含まれる
- log_formatディレクティブ : ログのフォーマット
- access_logディレクティブ : ログの出力先

### log_formatディレクティブ
- \t : tab文字
- 常に任意のフォーマットの設定をインクルードしている

| log_format | 出力するログの書式を定義する                              |
| ---------- | --------------------------------------------------------- |
| 構文       | log_format フォーマット名 ログの書式文字列 ...;           |
| default    | combiner '$remote_addr - $remote_user [$time_local '] ... |
| context    | http                                                      |


### access_log ディレクティブ
- offを指定すると出力なし

| access_log | 出力するログの書式を定義する                         |
| ---------- | ---------------------------------------------------- |
| 構文       | access_log ファイルパス [フォーマット];              |
| default    | logs/access.log combined                             |
| context    | http, server, location, location中のif, limit_except |


# 3.3 nginx本体の設定
nginx本体の基本的な設定
```
user nobody;
worker_process 1;

error_log /var/log/nginx/error.log;
pid /var/run/nginx.pid;

events {
    worker_connections 1024
}

http {
    server {
        listen 80;
    }
}
```


## エラーログの出力設定

### error_log ディレクティブ
- 設定されているエラーレベル以上のエラーを出力

| error_log | 出力するエラーログのファイルパスを定義する |
| --------- | ------------------------------------------ |
| 構文      | error_log ファイルパス [エラーレベル];     |
| default   | logs/error.log error                       |
| context   | main, http, server, location               |

### バーチャルサーバ別のエラーログファイル指定
- mainコンテキスト二記述した場合nginx本体に関するすべてのエラーが出力
- serverディレクティブ内に記述した場合、serverコンテキストにマッチするリクエストのエラーのみが出力される

### log_not_found ディレクティブ
- 要求されたファイルが存在しない場合にエラーを出力するか否かを決定

| log_not_found | 要求ファイルが存在しない場合のエラーを出力するか否か |
| ------------- | ---------------------------------------------------- |
| 構文          | log_not_found on or off;                             |
| default       | on                                                   |
| context       | http, server, location                               |


## プロセスの動作に関する設定
- masterプロセスの管理に必要なPIDファイルの出力先の指定
- ワーカプロセスの動作に関する設定

### pid ディレクティブ
- PIDファイルの出力先
- default : logs/nginx.pid

### user ディレクティブ
- ワーカプロセスの実行ユーザを指定する
- default : nobodyユーザ
- debianではhttpサーバをwww-dataユーザで動作させる

### worker_processes ディレクティブ
- ワーカのプロセス数の指定
- worker_processes = auto の時、CPUのコア数と同じ数が動作する

### worker_rlimit_nofile ディレクティブ
- 1プロセスで同時にオープン可能なファイルディスクリプタの数を指定する
- OSによって一度に開くことのできるファイルディスクリプタの数は決まっている
  - この値を越えないように設定？

### events ディレクティブ
- ワーカのイベント駆動方式に関連するディレクティブを記述するディレクティブ
- このディレクティブは省略できない（内容が空でもかく）


### worker_connectinos ディレクティブ
- workerが処理するコネクション数を指定する
- workerごとの最大数
- エラー : worker_connections are not enough

### use ディレクティブ
- 利用するコネクションの処理方式
- 通常は必要ない


# 3.4 パフォーマンスに影響する設定

## keepalive_timeout ディレクティブ
- nginxに常時接続しているクライアントに対するタイムアウト時間を設定する\
- 第二パラメータにKeep-Aliveヘッダに付加するタイムアウト時間を付加できる

## sendfile ディレクティブ
- sendfile()システムコールを使用するか否か
- 基本的には効率よくなる

## tcp_nopush ディレクティブ
- on にするとパケットサイズをなるべく大きくして、送信するパケット数を少なくすることができる。

## open_file_cache ディレクティブ
- on の時、一度オープンしたファイルの以下の情報をキャッシュする
  - ファイルのディスクリプタ、サイズ、更新日時
  - ディレクトリが存在するか
  - ファイルが存在しない、読み取り権限エラーなどの情報
- max : キャッシュ可能な数
- inactive : 有効期限
- エラーをキャッシュするには、別途open_file_cache_errorsを有効化する必要がある

## worker_cpu_affinity ディレクティブ
- nginxの各ワーカがどのCPUに割り当てられるかを規定するもの
- 必要となるケースは少ない

## pcre_jit ディレクティブ
- jitコンパイルができるようになる

