// グーグルにリダイレクトするヘッダの書き込み
// 1.curl -i 127.0.0.1:8080/redirect
// 2.httpレスポンスを確認
// 3.実際にブラウザ上で確認

package main

import (
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello Osaka</h1></body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。あほんだらピーポー")
}

func HeaderExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", HeaderExample)
	server.ListenAndServe()
}
