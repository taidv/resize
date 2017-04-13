# Resize

画像のサイズ変更と圧縮ツール<br />
https://github.com/nfnt/resize の実装<br />

## 環境準備

### Go言語のインストール
https://golang.org/doc/install<br />
http://golang.jp/install<br />

### resizeをゲット
```
$ go get github.com/nfnt/resize
```

### GOプログラム作成
```
$ cp -r ~/resize $HOME/go/src/resize
$ cd $HOME/go/src/resize
$ go install
```
## パラメータ設定
### パス設定
`resize.go`ファイルの17,18行にパスを変更ください。<br />

`rawDir`          入力ディレクトリ<br />
`resizedDir`      出力ディレクトリ<br />
<br />
例
```
var rawDir string = "~/Downloads/media/"
var resizedDir string = "~/Desktop/media/"
```
### 圧縮率設定
`resize.go`ファイルの16行に圧縮率を変更ください。<br />
`quality`         範囲は1から100まで。<br />
<br />
例
```
const quality = 50
```
設定が終わった際に以下のコマンドを実行ください。
```
$ go install
```

## 使い方

```
$ resize
```
