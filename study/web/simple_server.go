package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	// %q 会添加双引号
	fmt.Fprintf(w, "Hello url.path = %q\n", r.URL.Path)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	mu.Unlock()
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// 获取 form 表单内容
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "form: [%q] = %q\n", k, v)
	}
}
