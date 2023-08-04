## Regex Searcher

![regex_searcher_logo](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/cbd14194-c790-4ffe-9401-48e418fc3bfd)

このプログラムは、対象のディレクトリ内に特定のパターンが存在するかどうかを判定するためのものです。

対象のディレクトリとは、`target_files`ディレクトリです。

特定のパターンとは、`数値 + 非数値文字列`です。
（詳細は後述します）

## デモ

![Aug-04-2023 14-42-10](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/2c49f823-60b5-4d68-89a2-864fc1343021)

## 使い方

以下の手順で使用できます。

### 手順1. クローン

通常通り、ローカル環境へコードをクローンしてください。

### 手順2. コンテナ起動

プロジェクトのルートディレクトリで、以下のコマンドを実行します。

```sh:コマンドラインツール
$ cd regex_searcher

$ docker compose up -d
```

### 手順3. サンプルファイルを使った処理実行

プロジェクトのルートディレクトリで、以下のコマンドを実行します。

```sh:regex_searcher
$ docker compose exec -it regex_searcher go run main.go
```

以下のように動作すれば成功です。

![Aug-04-2023 14-42-10](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/2c49f823-60b5-4d68-89a2-864fc1343021)

失敗する場合は、エラーメッセージを見ながら対処してください。

### 手順4. 処理対象ファイルを使った処理実行

特定のパターンが存在するかを調べたいファイルを、`target_files`ディレクトリに格納します。

その後、手順3. を再度実行してください。

対象文字列が存在する場合、以下のように表示されます。

```
対象文字列が存在します: (ファイルパス)
  (行数): (対象文字列が存在する箇所の記述)
```

例えば、以下のように表示されます。
```
対象文字列が存在します: target_files/sample_match.txt
  L1: 123abc`
```

対象文字列が存在しない場合、以下のように表示されます。

`対象文字列は存在しません: (ファイルパス)`
