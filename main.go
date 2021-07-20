package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

// Go言語においてハンドラ（Handler）とは、
// ServeHTTPというメソッドを持ったインタフェースのことです。
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type WorldHandler struct{}

func (wh *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

// ハンドラとはメソッドServeHTTPを持つインターフェース
// ハンドラ関数は、ハンドラのように振る舞う関数
// Handlerのほうが拡張性が高い

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
