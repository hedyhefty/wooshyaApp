package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func CpyIndex(res http.ResponseWriter, req *http.Request) {

	// fmt.Println("PPath: ", PPath)
	// fmt.Println(PPath + "/views/index.html")

	//add by st
	fmt.Println("call Chp.")
	cpyHptpl, err := template.ParseFiles(PPath+"/views/cpyIndex.html", PPath+"/views/hnavbartpl.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		panic(err.Error())
		return
	}

	var cpyHptplHandler Indextpl

	cookies, err := req.Cookie("SessionID")
	if err != nil {
		fmt.Println("No cookies.")
		cpyHptplHandler.IsOnline = false
	} else {
		fmt.Printf("cookies: %s = %s\n", cookies.Name, cookies.Value)
		if SessionMap[cookies.Value] != nil {
			session := SessionMap[cookies.Value]
			if session.SessionType == Company {
				cpyHptplHandler.Username = session.Username
				cpyHptplHandler.IsOnline = session.Connection

				fmt.Printf("User name is: %s\nSession type is: %d\nStatus of connection is: %t\n", session.Username, session.SessionType, session.Connection)

			}
		} else {
			cpyHptplHandler.IsOnline = false
			fmt.Println("Unknown cookies value.")
		}
	}

	err = cpyHptpl.Execute(res, cpyHptplHandler)
	if err != nil {
		panic(err.Error())
		return
	}
	//

	//http.ServeFile(res, req, PPath+"/views/companyIndex.html")
}
