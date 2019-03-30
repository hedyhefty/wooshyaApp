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
			SessionMap[cookie.Value].SetOff()

			fmt.Println("Status of Session: ",SessionMap[cookie.Value].Connection)
		}
	}

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
