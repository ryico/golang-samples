package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //オプションを解析する。デフォルトでは解析しない。
	fmt.Println(r.Form) //このデータはサーバのプリント情報に出力される
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //ここでwに入るものがクライアントに出力される
}

func main() {
	http.HandleFunc("/", sayhelloName) //アクセスのルーティングを設定
	err := http.ListenAndServe(":9090", nil) //監視するポートを設定
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}