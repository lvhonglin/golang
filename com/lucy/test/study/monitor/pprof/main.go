package main

import (
	"net/http"
	"net/http/pprof"
)

func main() {
	mux := http.NewServeMux()
	//注册prof处理器
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	//启动http服务器
	http.ListenAndServe(":8080", mux)
}
