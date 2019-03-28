package routerHandlers

import (
	"fmt"
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {

	//fmt.Println("PPath: ", PPath)
	//fmt.Println(PPath + "/views/index.html")

	cookie, err := req.Cookie("username")
	if err != nil {
		fmt.Println("No cookies.")
	} else {
		fmt.Printf("%s=%s\r\n", cookie.Name, cookie.Value)
		if SessionMap[cookie.Value] != nil {
			session := *SessionMap[cookie.Value]
			var status string
			if session.Connection {
				status = "Online"
			} else {
				status = "Offline"
			}
			fmt.Printf("Session status of %s is: %s\n", session.Username, status)
		} else {
			session := Session{cookie.Value, false}
			SessionMap[cookie.Value] = &session
		}
	}

	//connecting := SessionMap[cookie.Value]
	//if connecting != nil {
	//	fmt.Printf("Session %s 's connection is: %t", (*connecting).Username, (*connecting).Connection)
	//}

	http.ServeFile(res, req, PPath+"/views/index.html")
}
