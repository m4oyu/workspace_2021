# 過去問 ISUCON 9
isucon9の過去問を説いた時のメモ

## 環境設定など
- amazon ec2 ssh
  - ssh -i path/file-name.pem instance-os-name@instance-dns-name
  - example^ : ssh -i ./example.pem ubuntu@ec2-1-1-1-1.ap-northeast-1.compute.amazonaws.com
  - インスタンスのセキュリティ設定で自分のIPアドレスからのみアクセスできるようにしていたが、時間がたって自分のIPアドレスが変わり、アクセスできなくなった
- github releases assetsのダウンロード
  - curlコマンドでダウンロード用のURLをたたけばよい
- cat example.sql | sudo mysql とかでsqlファイルを走らせることができた


## linuxコマンド類
- stat [filename] : ファイルの詳細情報
- rm -rf [filename] : ファイルの削除
  - -r : ディレクトリごと削除
  - -f : エラーメッセージを表示せず削除
- ls -d */ : ディレクトリのみ出力
- ls | grep [regex] : これで好きなファイルを見れる
- curl [url] : 様々なプロトコルでファイルをダウンロードする
  - -s : ダウンロード中の出力を表示しない
  - -L : リダイレクトに従う
  - -J : http-headerが提供するファイル名で保存する
  - -O : 転送元と同じ名前で保存する