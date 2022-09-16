# goによるCRUD処理API
goの学習過程で簡易的なAPIを作成しました。
DBと本のデータをやりとりします。

## 使用技術/動作環境
* go
* XAMPP 8.1.6
* Apache/2.4.53
* MariaDB 10.4.24
* Postman

## Postmanを利用して動作確認

### 1: Get(http://localhost:8080/posts)
あらかじめ用意した２つのデータを取得。![スクリーンショット 2022-09-16 135819](https://user-images.githubusercontent.com/106419953/190563596-b3dfdb5f-0d72-446e-ab79-3624ea880517.png)

### 2: Get(http://localhost:8080/posts/1)
id=1のデータを取得。![スクリーンショット 2022-09-16 135850](https://user-images.githubusercontent.com/106419953/190563828-49976f55-61a9-497c-8d7c-eb40fee15ba2.png)

### 3: Post(http://localhost:8080/posts/)
データの新規作成。![スクリーンショット 2022-09-16 140122](https://user-images.githubusercontent.com/106419953/190564019-3dc541dc-5cea-4ab5-bb29-a4ce6d69be04.png)

### 4: Put(http://localhost:8080/posts/)
データの更新
![スクリーンショット 2022-09-16 140331](https://user-images.githubusercontent.com/106419953/190564271-199b9005-6c91-4dc0-a80a-6fba8bc6fbef.png)

### 5: DELETE(http://localhost:8080/posts/)
データの削除![スクリーンショット 2022-09-16 140400](https://user-images.githubusercontent.com/106419953/190564377-71eb8cc3-0e04-4e36-b755-8257b43a3102.png)

## 参考
以下の資料、サイトを参考にさせていただきました。ありがとうございました。
* https://www.chuken-engineer.com/entry/2021/09/24/162120
* サウション・チャン, 武舎広幸, 阿部和也, 上西昌弘 『Goプログラミング実践入門　標準ライブラリでゼロからWebアプリを作る』 インプレス 2017
