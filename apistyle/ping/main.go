package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pong)
	log.Println("Starting http server ...")
	log.Fatal(http.ListenAndServe(":50021", nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

// 测试
// $ curl http://127.0.0.1:50021/ping
// pong
