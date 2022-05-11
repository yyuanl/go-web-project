package router

import (
	"net/http"
	"web_project/controller"
)

func Router() {

	http.HandleFunc("/", controller.IndexPageHandle)
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer,"url.path=%q\n",request.URL.Path)
	//})
	http.HandleFunc("/loginPage", controller.LoginPageHandle)

	http.HandleFunc("/registerPage", controller.RegisterPageHandle)
	http.HandleFunc("/login", controller.LoginHandle)

	http.HandleFunc("/register", controller.RegisterHandle)

	http.HandleFunc("/checkUserName", controller.CheckUserIsExistHandle)
}
