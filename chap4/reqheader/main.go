package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}

// map[
// Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9]
// Accept-Encoding:[gzip, deflate, br]
// Accept-Language:[ja,en-US;q=0.9,en;q=0.8]
// Cache-Control:[max-age=0]
// Connection:[keep-alive]
// Sec-Ch-Ua:["Chromium";v="92", " Not A;Brand";v="99", "Google Chrome";v="92"]
// Sec-Ch-Ua-Mobile:[?0]
// Sec-Fetch-Dest:[document]
// Sec-Fetch-Mode:[navigate]
// Sec-Fetch-Site:[none]
// Sec-Fetch-User:[?1]
// Upgrade-Insecure-Requests:[1]
// User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36]
// ]
