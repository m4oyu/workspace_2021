# Go System Programing
書籍 :  [Goならわかるシステムプログラミング](https://www.amazon.co.jp/Go%E3%81%AA%E3%82%89%E3%82%8F%E3%81%8B%E3%82%8B%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0-%E6%B8%8B%E5%B7%9D-%E3%82%88%E3%81%97%E3%81%8D/dp/4908686033/ref=cm_cr_arp_d_product_top?ie=UTF8)


### chap 3 memo
3.4: 標準入力、ファイル入力、ネットワーク通信、メモリ読み出しバッファをio.Readerインタフェースで共通に扱う\
3.5: バイナリ解析\
3.6: テキスト解析、Scanner, fmt.Fscanf, csv, json\
3.7: I/O多重化\

### chap 4 memo
channelで並列処理がとても簡単にできる。goroutineで必要な時に必要なデータを取り出す処理、並行して処理を行い、メインルーチンに通知するなど


### chap 6 memorandom
tcpとHTTPによるソケット通信についての内容。HTTP1についての基礎的な学習でHTTP通信とはどんなものなのかを学習して、その高速化手法としてKeep-Aliveや並列処理などを行った。どのようなプロセスを経てHTTP2へと変化したのかについての内容もあり勉強になった。