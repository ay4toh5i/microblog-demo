# microblog demo

## how to develop
```sh
$ docker compose up
```

apiサーバーが`localhost:8000`, フロントエンドのdevサーバーが`localhost:3000`で立ち上がります。  
apiサーバーが立ち上がる際にマイグレーションが走りますが、CLIでの実行も可能なので、ローカルではCLIを使うと便利です。  
[golang-migrate/migrate](https://github.com/golang-migrate/migrate)

## Directory structure

### backend
```
backend
├── app
│   └── web       httpのアプリケーション層
├── cmd
│   └── api       実行用のエンドポイント
├── config
│   └── config.go 
├── db
│   └── mysql     マイグレーションファイルなど
└── microblog
    ├── domain    主なドメインに関する処理
    ├── infra     repository層の実装
    └── usecase   各ユースケースに対するインターフェース
```

### frontend
```
frontend
└── src
    ├── components 
    ├── context    
    ├── http      HTTPクライアントのカスタムインスタンス
    ├── pages     
    ├── styles    全体で共通のスタイルやテーマをおく
    └── types     共通の型定義やライブラリの定義の上書きなど
```

## TODO
- トランザクションの実装
- DIの整理
- バリデーションやテストの追加
