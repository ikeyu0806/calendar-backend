# calendar-backend
[vue.jsで実装したカレンダー](https://github.com/ikeyu0806/vue-calendar)のバックエンドです。

APIインターフェイスにはGraphQLを使用しています。

# Migration実行
`sql-migrate up`

# GraphQL
`graph/schema.graphqls`記述後に`gqlgen generate`コマンドでソースコード生成
