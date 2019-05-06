package controllers

import (
	"fmt"
	"net/http"
)

func CpyLogOut(w http.ResponseWriter, r *http.Request) {
	cookies, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println("No cookies.")
		http.Redirect(w, r, "/cpyIndex", http.StatusSeeOther)
		return
	}

	if SessionMap[cookies.Value] != nil {
		session := SessionMap[cookies.Value]
		if session.SessionType == Company {
			session.SetOff()
		} else {
			fmt.Println("Session ID not matching.")
		}

	} else {
		fmt.Println("Session not exist.")
	}

	http.Redirect(w, r, "/cpyIndex", http.StatusSeeOther)
}
