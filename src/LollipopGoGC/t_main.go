package main

import (
	//"log"
	//	"net"
	"glog-master"
	"net/http"

	//	"golang.org/x/net/context"
	//	"google.golang.org/grpc"
	//	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	//	"google.golang.org/grpc/reflection"
)

// 接受数据处理
func TJWanJiaData(w http.ResponseWriter, req *http.Request) {
	glog.Info("httpTask is running...")
	if req.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()
		// 获取函数
		Protocol, bProtocol := req.Form["Protocol"]
		Protocol2, bProtocol2 := req.Form["Protocol2"]
		glog.Info("httpTask is running...", Protocol, bProtocol, Protocol2, bProtocol2)
	}
}

// 主函数
func main() {
	http.HandleFunc("/TJData", TJWanJiaData) // 获取统计
	err := http.ListenAndServe(":7878", nil)
	if err != nil {
		glog.Info("Entry nil", err.Error())
		return
	}
}
