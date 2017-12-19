# translation-tool
コマンドラインで翻訳するツール。コーディング時の変数名決めるのに使ったりする

# install 

```bash
go get github.com/ahaha0807/translation-tool/cmd/trans
```

# usage

使う前に Azureでtranslator apiを使えるようにし、そのAPIキーを 環境変数 `MICROSOFT_TRANSLATE_APIKEY` に登録してください。

```bash
trans [日本語(Japanese word)]

trans -e [英語(English word)]
```

