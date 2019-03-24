package routerHandlers

import (
	"fmt"
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
	err := DB.Ping()
	if err != nil {
		panic(err.Error())
	}else {
		println("I know u DB.")
	}
	fmt.Println("PPath: ", PPath)
	fmt.Println(PPath + "/views/index.html")
	http.ServeFile(res, req, PPath+"/views/index.html")
}

//func SignUpPage(res http.ResponseWriter, req *http.Request){
//	if (*req).Method != "POST"{
//		http.ServeFile(res,req,"signup.html")
//		return
//	}
//
//	username := req.FormValue("username")
//	password := req.FormValue("password")
//	mailaddress := req.FormValue("mailaddress")
//	collegename := req.FormValue("collegename")
//	degree := req.FormValue("degree")
//	department := req.FormValue("department")
//	major := req.FormValue("major")
//
//}