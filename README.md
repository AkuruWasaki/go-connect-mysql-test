# go-connect-mysql-test
本アプリはgolang + gin 学習用のアプリとなります。

# 機能
- ユーザー情報一覧
- ユーザー作成
- ユーザー情報更新
- ユーザー削除

# これからの課題
- ORMにGORMではなくSQLBoilerを使いたかった。。
  - ```sqlboiler mysql``` コマンドでDBからモデルを生成しようとしたが、次のエラーにある通り、 DBとの接続がうまくいかなかった
  
  ```
  unable to get table names: dial tcp 127.0.0.1:3306: connect: connection refusedError: unable to initialize tables: unable to fetch table data: driver (/go/bin/sqlboiler-mysql) exited non-zero: exit status 1
  ```
  - 原因については調査中
  - わかっていること
    - 同じUSER, PASSWORD, Host, Portの入力情報でGORM使用時とSQLクライアント使用時での接続は確認できた。
- API化して動かしてみる
- Azure(クラウド環境)にデプロイしてみる。

# 環境構築手順

1. git cloneし、ソースコードをローカルへ落とす
```
git clone https://github.com/AkuruWasaki/go-connect-mysql-test
```

2. go.modファイルの作成
```
go mod init go-connect-mysql-test
go mod tidy
```

3. コンテナのbuildと実行
```
docker-compose up --build
```

4. ブラウザでローカル環境へアクセス
```
http://localhost:8080
```
