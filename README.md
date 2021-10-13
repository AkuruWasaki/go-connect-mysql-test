# go-connect-mysql-test
本アプリはgolang + gin 学習用のアプリとなります。

# 機能
- ユーザー情報一覧
- ユーザー作成
- ユーザー情報更新
- ユーザー削除

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

3. コンテナ作成
```
docker-compose up --build
```

4. ブラウザでローカル環境へアクセス
```
http://localhost:8080
```
