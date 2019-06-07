// ブラウザへのクッキーの送信
// 1.ブラウザのURLの鍵をクリック
// 2.中身を確認

package main

import (
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func main() {
	sever := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	sever.ListenAndServe()
}
