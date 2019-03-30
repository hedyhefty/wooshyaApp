package routerHandlers

import (
	"html/template"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CompanyLogin(res http.ResponseWriter, req *http.Request) {
	if (*req).Method != "POST" {

		//by st
		cpylogintpl, err := template.ParseFiles(PPath+"/views/companyLogin.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			panic(err.Error())
			return
		}

		err = cpylogintpl.Execute(res, nil)
		if err != nil {
			panic(err.Error())
			return
		}
		//

		//http.ServeFile(res, req, PPath+"/views/companyLogin.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databasePassword string

	err := DB.QueryRow("select password from companyuser where username = ?", username).Scan(&databasePassword)
	if err != nil {
		// todo: add a info box including "Username not existed"
		http.Redirect(res, req, "/companyLogin", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		// todo: add a info box including "Wrong password"
		http.Redirect(res, req, "/companyLogin", 301)
	}

	updateDatehandler, err := DB.Prepare("UPDATE companyuser SET lastlogindate = ? WHERE username = ?")
	if err != nil {
		// todo: add a info box to inform loginDate error
		http.Redirect(res, req, "/companyLogin", 301)
		return
	}

	logindate := time.Now().Local()
	_, err = updateDatehandler.Exec(logindate, username)
	if err != nil {
		http.Redirect(res, req, "/companyLogin", 301)
		return
	}

	session := Session{Username: username, Connection: true, SessionType: Company}
	session.StartSession()
	go session.EndSession()

	cookies := session.SetCookies()
	http.SetCookie(res, &cookies)

	http.Redirect(res, req, "/companyIndex", http.StatusSeeOther)
}
