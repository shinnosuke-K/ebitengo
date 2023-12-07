# ebitengo

Go製の2Dゲーム用の [Ebitengine](https://ebitengine.org/) の使ったゲームを実装しています。

[eihigh/wasmgame](https://github.com/eihigh/wasmgame) をテンプレートとして使用して、編集をしています。

https://shinnosuke-k.github.io/ebitengo で実装したゲームをプレイできます。

スマートフォンからは操作することはできないと思うので、PCから操作をお願いします。


### 実行方法
```shell
# 実行後、ブラウザで http://localhost:8080 にアクセス
$ go run ./tool serve

# 別ターミナルで実行（以後、このコマンドを実行するとファイルが更新される）
$ go run ./tool build
```