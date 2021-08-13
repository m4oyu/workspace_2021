# Go Test 入門
参考：https://future-architect.github.io/articles/20200601/

### memo
任意のテスト実行
```shell
go test -run Test???
```

## testの失敗をレポート
- T.Error(T.Errorf)はアサーションのみ
- T.Fatal(T.Fatalf)はアサーションし、以降のテストを無視する
    - テスト失敗時に以降のテストが無意味になる場合に使う
    
## testのスキップ
- go test -v -shortとshortフラグを付けることでテストの自動スキップができる

## testを並列にしたい
- test.Parallel()により並列化
- tt := ttとしてtest用変数の補足を行う
- DBを伴う処理など並列化のできないテストも存在する

## testの前処理や後処理
- TestMain関数を使う
- TestMainを書いていれば、TestXXXが直接実行されずにTestMain内m.Run()で実行するので、前後に処理を挟める
- 同一パッケージ内に1回しか定義できない

## testのカバレッジを取得する
- -couvermode=count：カバレッジを取得
- -couverprofile=c.out：結果をファイルに出力


```shell
go test io/... -covermode=count -coverprofile=c.out
```

## あるディレクトリ配下のテストをすべて実施する
- ...の文字列はワイルドカード

```shell
go test io/...
```

## 一部のテストケースのみ実施したい
- -run [検索内容] で一致するテストのみを実行できる
- 例：Pipeにが含まれるテスト
```shell
go test -v io/... -run Pipe
```

## サブテストのみ実行
- これも-runフラグを使う
- 例：test関数にAddを含む＆サブテスト名にmalを含む testのみ実行
```shell
go test -v -run Add/mal
```

## テストのキャッシュを削除
- -count=1とすればキャッシュされない
- go clean -cache でビルドのキャッシュと共にテストのキャッシュの削除がされる

## テストのひな形を簡単に作成したい
- IDEで簡単にできる（個人的にできてないがとばす

## 構造体、マップやスライスの比較を実施する
- reflect.DeepEqualでこれらの値を比較できる
- github/google/go-cmpはwantとgetで何が違うのかきれいに出力してくれる


## APIサーバにアクセスするテスト
- net/http/httptestを用いれば簡単にテスト用のモックサーバがたてれる
- googleAPIなどをたたく場合にAPIサーバをモックできる
  - googleAPIにアクセスするためのクライアント構造体などのアクセス先を張り替えてテストを行う
  
## APIサーバハンドラのテスト
- httptest.NewRequest()で作られたリクエストでEchoやgo-swagger用のAPIハンドラもテストできる（らしい
  - ハンドラの単体テストではhttptest.NewRequestを用いることになる
