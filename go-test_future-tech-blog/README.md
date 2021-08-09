# Go Test 入門
参考：https://future-architect.github.io/articles/20200601/

## testの失敗をサポート
- T.Error(T.Errorf)はアサーションのみ
- T.Fatal(T.Fatalf)はアサーションし、以降のテストを無視する
    - テスト失敗時に以降のテストが無意味になる場合に使う
    
