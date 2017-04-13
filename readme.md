# Resize

画像のサイズ変更と圧縮ツール
https://github.com/nfnt/resizeの実装

## 環境準備

###Go言語のインストール
https://golang.org/doc/install
http://golang.jp/install

### resizeをゲット
```
$ go get github.com/nfnt/resize
```

### ソースをクローン
```
$ cd $HOME/go/src
$ git clone https://github.com/taidv/resize
```
## パラメータ設定
### パス設定
`resize.go`ファイルの17,18行にパスを変更ください。

`rawDir`          入力ディレクトリ
`resizedDir`      出力ディレクトリ

例
```
var rawDir string = "/Users/nci/Downloads/media/"
var resizedDir string = "/Users/nci/Desktop/media/"
```
### 圧縮率設定
`resize.go`ファイルの16行に圧縮率を変更ください。
`quality`         範囲は1から100まで

例
```
const quality = 50
```
設定が終わった際に以下のコマンドを実行ください。
```
go install
```

## 使い方

```
$ resize
```


