package main

import (
	"log"
	"net/http"
	"web_project/config"
	"web_project/dao"
	"web_project/router"
)

func main() {
	config.Init()

	dao.InitDb()
	httpRegister()
}

func httpRegister() {
	fs := http.FileServer(http.Dir("./views/static")) // 文件系统处理器
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fs1 := http.FileServer(http.Dir("./views/pages")) // 文件系统处理器
	http.Handle("/pages/", http.StripPrefix("/pages/", fs1))

	router.Router()

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalf("http服务错误：%+v", err)
	}
}
