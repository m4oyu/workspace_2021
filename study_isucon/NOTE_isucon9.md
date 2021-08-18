# 過去問 ISUCON 9
isucon9の過去問を解いた時のメモ

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

## bug
- isucariのアプリがブラウザから見れない件

  AWSのポート開放で0~65535を設定していない

- privateなremoteリポジトリに対して git push -u origin main するとエラーを吐く
  - sshkeyで認証すべき?

    $ ~/.sshで以下を実行
    ```
    ssh-keygen -t rsa
    ```
    参考\
    https://programming-juku.net/github-ssh-nopassword
    https://www.virtual-surfer.com/entry/2018/04/13/190000

  - gitのversionが古い？

    ubuntuなどではgithubのソースコードを持ってくる方が新しいのを取ってきやすい
    ```
    $ sudo apt -y install libcurl4-gnutls-dev libexpat1-dev gettext libz-dev libssl-dev autoconf asciidoc xmlto docbook2x make gcc
    $ wget https://github.com/git/git/archive/refs/tags/v2.33.0.tar.gz
    $ tar-zxf v2.33.0.tar.gz
    $ cd git-2.33.0
    $ make configure
    $ ./configure --prefix=/usr
    $ make all doc info
    $ sudo make install install-doc install-html install-info
    ```
    参考：https://qiita.com/noraworld/items/8546c44d1ec6d739493f

  - gitで管理する量が多すぎる件
    
    gitで管理できるのは100MBが上限\
    実装が分かれてるディレクトリのgoのディレクトリ以下のみを管理するべき

  - gitのdefaultブランチがmasterになってる件

    ~/.gitconfigを以下のように変更
    ```
    [init]
      defaultBranch = main
    ```

  - push先のURLに問題がある
    ```cmd
    ERROR: Repository not found.
    fatal: Could not read from remote repository
    ```
    git remote -v の値を以下のように変えた
    ```
    // fail
    git@github.com:user/repo.git/
    // correct
    git@github.com:user/repo.git
    ```