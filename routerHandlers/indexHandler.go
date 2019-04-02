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

func StdIndex(res http.ResponseWriter, req *http.Request) {

	//fmt.Println("PPath: ", PPath)
	//fmt.Println(PPath + "/views/index.html")

	//deal with the request to favcion, which missing until now.
	if req.URL.Path == "/favicon.ico" {
		return
	}

	fmt.Println("call Homepage.")

	indextplhandler, err := template.ParseFiles(PPath+"/views/index.html", PPath+"/views/navbartpl.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		fmt.Println("failed to load template.")
		return
	}

	var indextpl Indextpl

	cookie, err := req.Cookie("SessionID")

	if err != nil {
		fmt.Println("No cookies.")
		indextpl.IsOnline = false
	} else {
		fmt.Printf("%s=%s\r\n", cookie.Name, cookie.Value)
		if SessionMap[cookie.Value] != nil {
			session := SessionMap[cookie.Value]
			if session.SessionType == Student {
				var status string
				//Set username
				indextpl.Username = session.Username

				fmt.Println("username: ", indextpl.Username)
				fmt.Println("SessionType: ", session.SessionType)
				if session.Connection {
					status = "Online"
					indextpl.IsOnline = true
				} else {
					status = "Offline"
					indextpl.IsOnline = false
				}
				fmt.Printf("Session status of %s is: %s\n", session.Username, status)
			}
		} else {
			indextpl.IsOnline = false
			fmt.Println("Unknown Cookies value.")
			//indextpl.Username = ""
			//session := Session{cookie.Value, false}
			//SessionMap[cookie.Value] = &session
			//fmt.Printf("Create session for %s\n", cookie.Value)

		}
	}

	err = indextplhandler.Execute(res, indextpl)
	if err != nil {
		panic(err.Error())
		return
	}
}
