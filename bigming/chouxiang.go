package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	// pattern路由地址
	// handleFunc 处理方法
	Route(pattern string, handleFunc http.HandlerFunc)
	// address 监听的端口号
	Start(address string)
}

// 首字母大写的标识符是导出的（public），可在包外部访问；
// 首字母小写的标识符是非导出的（private），只能在当前包内访问。
type sdkHttpServer struct {
	Name string
}
type Header map[string][]string

// var headers = Header{
// 	"name": {
// 		"aaa", "bbbb",
// 	},
// }
// headers["name"] = []string{"ccc", "ddd"}

func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	// panic("出错了")
	http.HandleFunc(pattern, handleFunc)
}
func (s *sdkHttpServer) Start(address string) {
	// panic("出错了")
	http.ListenAndServe(address, nil)
}
func NewHttpServer() Server {
	return &sdkHttpServer{} // 必须实现了Server中的方法
}
func main() {
	http.HandleFunc("/coder", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("路由进入~")
	})
	http.ListenAndServe("localhost:8099", nil)
}
