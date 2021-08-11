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

## testの並列化
- test.Parallel()により並列化
- tt := ttとしてtest用変数の補足を行う
- DBを伴う処理など並列化のできないテストも存在する

## testの前処理や後処理
- TestMain関数を使う