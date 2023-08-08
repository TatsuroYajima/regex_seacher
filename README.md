## Regex Searcher

![regex_searcher_logo](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/cbd14194-c790-4ffe-9401-48e418fc3bfd)

このプログラムは、`target_files`ディレクトリ内に特定のパターンが存在するかどうかを判定できます。

特定のパターンとは、`数値 + 非数値文字列`です。
（詳細は [マッチする文字列](https://github.com/TatsuroYajima/regex_seacher#%E3%83%9E%E3%83%83%E3%83%81%E3%81%99%E3%82%8B%E6%96%87%E5%AD%97%E5%88%97)セクションを参照してください）

## デモ

![Aug-07-2023 17-45-09](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/f3830af2-b4d8-4787-95ca-8fc259183ace)

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

### 手順3. 処理対象ファイルを配置

特定のパターンが存在するか調べたいファイルを、`src/target_files`ディレクトリに格納します。

（試しに動かす場合は手順3はスキップ可能です）

![image](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/864adfbe-6c21-43ea-9128-2673f2d7ab79)

### 手順4. 処理実行

プロジェクトのルートディレクトリで、以下のコマンドを実行して下さい。

```sh:regex_searcher
$ docker compose exec -it regex_searcher go run main.go
```

以下のように動作すれば成功です。

![Aug-07-2023 17-45-09](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/f3830af2-b4d8-4787-95ca-8fc259183ace)

失敗する場合は、エラーメッセージを見ながら対処してください。

#### 対象文字列が存在する場合

結果ファイルに、以下のように出力されます。

```
対象文字列が存在します: (ファイルパス)
  (行数): (対象文字列が存在する箇所の記述)
```

例えば、以下のように出力されます。
```
対象文字列が存在します: src/target_files/sample_match.txt
  L1: 123abc`
```

#### 対象文字列が存在しない場合

結果ファイルに、以下のように出力されます。

`対象文字列は存在しません: (ファイルパス)`

## マッチする文字列

ファイル内に`数値 + 非数値文字列`が存在する場合は、正規表現にマッチします。

ただし、間にスペースが入っている場合（例：`123 abc`）はマッチしません。

|文字列|結果|
|---|---|
|`123abc`| マッチします|
|`select * from a where id = 123union all`| マッチします|
|`123`| マッチしません|
|`abc`| マッチしません|
|`123 abc`| マッチしません|

## テスト

テストを実行するには、以下のコマンドを実行します。

```sh
$ docker compose exec -it regex_searcher go test
```

テストに成功すると、以下のような表示になります。

![image](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/3f54cdd6-1adf-46ff-ac68-c45953729de4)

テストに失敗すると、以下のような表示になります。

![image](https://github.com/TatsuroYajima/regex_seacher/assets/97079648/21213734-bf8a-4788-bf07-f5a4876bd897)
