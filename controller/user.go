package controller

import (
	"fmt"
	"log"
	"net/http"
	"web_project/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("登录成功！"))
	userName := r.PostFormValue("username")
	passWord := r.PostFormValue("password")
	isExist, err := services.JudgeUserIsExist(userName, passWord)
	if err != nil {
		log.Printf("判断用户是否存在错误：%+v", err)
		SendHtmlTemplateData(loginHtmlPath, w, "判断用户是否存在错误")
		return
	}
	if isExist {
		w.Write([]byte("登录成功！"))
	} else {
		SendHtmlTemplateData(loginHtmlPath, w, "用户名或密码不正确")
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	// 获取post表单数据
	userName := r.PostFormValue("username")
	passWord := r.PostFormValue("password")

	fmt.Printf("用户名是：%s,密码是：%s，注册成功\n", userName, passWord)
	SendHtmlTemplateData(registerHtmlPath, w, "用户名已存在")
}
