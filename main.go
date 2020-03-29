package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	n := negroni.Classic()
	n.UseHandler(mux)
	hostString := fmt.Sprintf(":%s", port)
	n.Run(hostString)
}

//测试车服务器
//测试异步
//bus
//ces 测试构建模块 dev 测试本地环境构建 build 构建 deploy 部署 发布 Resgister
// 服务器部署方式
//需要在wercker进行工作流的设计即可 并行还是关系节点即可
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from Go!")
}
