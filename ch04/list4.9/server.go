// クライアントを転送するヘッダの書き込み
// 1.curl -i 127.0.0.1:8080/writeHeader
// 2.httpレスポンスを確認

package main

import (
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming<title></head>
	<body><h1>Hello World</h1><body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。他を当たって下さい。")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeHeaderExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	server.ListenAndServe()
}
