package main

import (
	"fmt"
	"log"
	"net/http"
)

// 这个过程显然比C++的socket编程要简洁一些

func main() {
	// 将路由与响应的具体操作进行绑定
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path %q\n", req.URL.Path)
}

// 打印返回的http的相应报文头部内容
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
