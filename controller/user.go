package controller

import (
	"log"
	"net/http"
	"web_project/services"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("登录成功！"))
	userName := r.PostFormValue("username")
	//passWord := r.PostFormValue("password")
	isExist, err := services.JudgeUserIsExist(userName)
	if err != nil {
		log.Printf("判断用户是否存在错误：%+v", err)
		SendHtmlTemplateData(loginHtmlPath, w, "判断用户是否存在错误")
		return
	}
	if isExist {
		w.Write([]byte("登录成功！")) // 应该展示登录成功页面
	} else {
		SendHtmlTemplateData(loginHtmlPath, w, "用户名或密码不正确")
	}

}

func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	// 获取post表单数据
	userName := r.PostFormValue("username")
	passWord := r.PostFormValue("password")

	isExist, err := services.JudgeUserIsExist(userName)
	if err != nil {
		log.Printf("判断用户是否存在错误：%+v", err)
		w.Write([]byte("服务端错误"))
		return
	}
	if isExist {
		SendHtmlTemplateData(registerHtmlPath, w, "用户名已存在")
	} else {
		err := services.CreateOneUserAccount(userName, passWord)
		if err != nil {
			w.Write([]byte("服务端错误"))
			return
		}
		w.Write([]byte("注册成功！"))
	}
}

// CheckUserIsExistHandle .
func CheckUserIsExistHandle(w http.ResponseWriter, r *http.Request) {
	userName := r.PostFormValue("username")

	isExit, err := services.JudgeUserIsExist(userName)
	if err != nil {
		log.Printf("判断用户是否存在错误：%+v", err)
		return
	}

	if isExit {
		sendStr(w, "用户名不可用")
	} else {
		sendStr(w, "<front style='color:green'>用户名可用</front>")
	}
}

// sendStr .
func sendStr(w http.ResponseWriter, data string) {
	w.Write([]byte(data))
}
