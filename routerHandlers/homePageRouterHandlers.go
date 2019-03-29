package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type Indextpl struct {
	Username string
	IsOnline bool
}

func HomePage(res http.ResponseWriter, req *http.Request) {

	//fmt.Println("PPath: ", PPath)
	//fmt.Println(PPath + "/views/index.html")

	cookie, err := req.Cookie("username")

	indextplhandler, err := template.ParseFiles(PPath+"/views/index.html", PPath+"/views/navbartpl.html", PPath+"/views/bootstrapHeader.html")

	var indextpl Indextpl

	if err != nil {
		fmt.Println("No cookies.")
	} else {
		fmt.Printf("%s=%s\r\n", cookie.Name, cookie.Value)
		if SessionMap[cookie.Value] != nil {
			session := *SessionMap[cookie.Value]
			var status string
			//Set username
			indextpl.Username = cookie.Value

			fmt.Println("username: ", indextpl.Username)
			if session.Connection {
				status = "Online"
				indextpl.IsOnline = true
			} else {
				status = "Offline"
				indextpl.IsOnline = false
			}
			fmt.Printf("Session status of %s is: %s\n", session.Username, status)
		} else {
			indextpl.IsOnline = false
			indextpl.Username = ""
			session := Session{cookie.Value, false}
			SessionMap[cookie.Value] = &session
			fmt.Printf("Create session for %s\n", cookie.Value)
		}
	}

	err = indextplhandler.Execute(res, indextpl)
	if err != nil {
		panic(err.Error())
		return
	}
}
