package routerHandlers

import (
	"fmt"
	"net/http"
)

func StdLogout(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("SessionID")

	if err != nil {
		fmt.Println("No cookies.")
	} else {
		fmt.Printf("%s=%s\r\n", cookie.Name, cookie.Value)
		if SessionMap[cookie.Value] != nil {
			session := SessionMap[cookie.Value]
			if session.SessionType == Student {
				session.SetOff()
			} else {
				fmt.Println("Session not matching")
			}
		} else {
			fmt.Println("Session not exist.")
		}
	}

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
