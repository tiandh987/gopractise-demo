package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", resp)
	log.Println("Start http server ...")
	log.Fatal(http.ListenAndServe(":50022", nil))
}

func resp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

// 测试
// $ curl http://127.0.0.1:50022/hello
// Hello World!
