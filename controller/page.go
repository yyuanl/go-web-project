package controller

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var (
	indexHtmlPath    = "./views/index.html"
	loginHtmlPath    = "./views/pages/users/login.html"
	registerHtmlPath = "./views/pages/users/register.html"
)

// IndexPageHandle 导航页
func IndexPageHandle(w http.ResponseWriter, r *http.Request) {
	sendHtmlTemplate(indexHtmlPath, w)
}
func TestParse() {
	_, err := template.ParseFiles(registerHtmlPath)
	fmt.Println(err)
}

// LoginPageHandle 登录页
func LoginPageHandle(w http.ResponseWriter, r *http.Request) {
	sendHtmlTemplate(loginHtmlPath, w)
}

// RegisterPageHandle 注册页
func RegisterPageHandle(w http.ResponseWriter, r *http.Request) {
	sendHtmlTemplate(registerHtmlPath, w)
}

// sendHtmlFile 原生的 读html文件 写入客户端
func sendHtmlFile(htmlPath string, w http.ResponseWriter) {
	f, err := os.OpenFile(htmlPath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("打开注册html文件错误:", err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	dataByte, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("读html文件错误:", err)
		return
	}

	io.WriteString(w, string(dataByte))
}

// sendHtmlTemplate 使用template
func sendHtmlTemplate(htmlPath string, w http.ResponseWriter) {
	t, err := template.ParseFiles(htmlPath)
	if err != nil {
		fmt.Printf("解析模板错误，htmlpath:%s,err:%+v\n", htmlPath, err)
	}
	//t := template.Must(template.ParseFiles(htmlPath))
	t.Execute(w, "")
}

func SendHtmlTemplateData(htmlPath string, w http.ResponseWriter, data string) {
	t := template.Must(template.ParseFiles(htmlPath))
	t.Execute(w, data)
}
