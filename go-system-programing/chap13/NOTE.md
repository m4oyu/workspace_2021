# chap 13 Go言語と並列処理

## スレッドとgoroutineの違い
Thread
- OSによって手配される
- メモリにロードされたプログラムの現在の実行状態を持つ仮想CPU
- 時間が凍結されたプログラムの実行状態
- 時間単位でCPUコアに割り当てられ、実行する

goroutine
- OSのスレッドに割り当てられる
- スレッドIDを持たない
- 初期スタックメモリサイズが小さく、起動が速い
- 優先度を持たない
- タイムスライスによる処理の中断がない
- 外部からの終了リクエストが飛んでこない

つまりはとても高速（OSのThreadの1000倍のオーダー）

## GoのランタイムはミニOS
goroutineのメリット
- サーバの実装にクライアントごとに一つのgoroutineを割り当てられる（メモリが小さいから）
- OSのThreadはスレッド切り替え時にコンテキストスイッチをする必要があるが、goroutineは切り替えが簡単
- deadlockのデバックが簡単

goのランタイム
- 内部でgoroutineのスケジューラを持っている
- 優先度を付けない
- goroutineの実行場所を指定できない

## 13.5 runtimeパッケージのgoroutine関連の機能
### 1 runtime.LockOSThread() / runtime.UnlockOSThread()
- runtime.LockOSThread()をコールすることで現在実行中のOSスレッドでのみgoroutineが実行されるように制限できる。
- そのスレッドが他のgoroutineによって使用されない
- runtime.UnlockOSThread()が呼ばれるか、goroutineが終了するかまでロックされる
- mainスレッドでの実行が強制されているライブラリの使用時に必要
- init関数は必ずmainスレッドで実行されるため、その時に固定する

### 2 runtime.Gosched()
現在進行中のgoroutineを一時中断して他のgoroutineに処理を回す

### 3 runtime.GOMAXPROCS(n) / runtime.NumCPU()
- 同時実行するOSスレッド数を制限する関数
- runtime.NumCPU()でCPU数がわかる

## 13.6 Race Detector
goroutineのデータ競合を発見する機能

## 13.7 syncパッケージ
他の言語で書かれたものをGoで置き換えるときに使うAPI

### 13.7.1 sync.Mutex / sync.RWMutex
- メモリの保護のための機能
- クリティカルセクションの同時実行によるデータの不整合を防ぐ


- チャネル: データの所有権を渡す場合、作業を並列化して分散する場合、非同期で結果を受け取る場合
- Mutex: キャッシュ、状態管理
    - sync.Mutex: 読み込みと書き込みが同時に行われる
    - sync.RWMutex: 複数のgoroutineで共有されるキャッシュの保護

### 13.7.2 sync.WaitGroup
- 多数のgoroutineで実行しているジョブの終了待ちを行う
- ジョブ数を数でカウントして０になるまでまつ

### 13.7.3 sync.Once
- 一度だけ呼び出したい処理に使用する
- 基本的にはinit関数で事足りる
- 初期化処理を必要な時まで遅延させたい場合に使用（initは最初に強制的によばれるため）

### 13.7.4 sync.Cond
- 条件変数と呼ばれる排他制御の仕組み
- ロック、アンロックでクリティカルセクションを保護する
  - 先に終わらせたいタスクがあり、完了したら待っているすべてのgoroutineに通知する（Broadcastメソッド）
      - TLUやHttp/2のハンドシェイク完了、レスポンスがきたタイミングでスレッドを起動する
  - リソースの準備が出来次第、そのリソースを待っているgoroutineに通知する(Signaiメソッド)
- waitでまち、Broadcastが来たら実行を再開する

### 13.7.5 sync.Map
- 大きなmapにアクセスする際にmap自体をロックするのは効率的でない
- ロックを内包し、複数のgoroutineからのアクセスに対応できるようになる


## sync/atomicパッケージ
- 不可分操作を提供するパッケージ
