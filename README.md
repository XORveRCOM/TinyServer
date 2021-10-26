# TinyServer

超手抜きのHTTPサーバ実装。
1. go build main.go でコンパイル
2. main.exe をルートディレクトリにコピーして実行
3. http://localhost:8080/ファイル名(index.htmlとか)
で、アクセスできます。

用途は、localStorage の作り込みで file:// アクセスはエラーになるため。
