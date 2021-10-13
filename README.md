# go-connect-mysql-test
本アプリはgolang + gin 学習用のアプリとなります。

# 機能
- ユーザー情報一覧
- ユーザー作成
- ユーザー情報更新
- ユーザー削除

# これからの課題
- ORMにGORMではなくSQLBoilerを使いたかった。。
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
