package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func StdLogin(res http.ResponseWriter, req *http.Request) {
	if (*req).Method != "POST" {
		logintpl, err := template.ParseFiles(PPath + "/views/stdLogin.html", PPath + "/views/bootstrapHeader.html")
		if err != nil{
			panic(err.Error())
			return
		}
		err = logintpl.Execute(res, nil)
		if err != nil {
			panic(err.Error())
			return
		}
		//http.ServeFile(res, req, PPath+"/views/stdLogin.html")
		return
	}


	username := GetFromValue(req,"username")
	fmt.Println(username)
	password := req.FormValue("password")

	var databasePassword string

	err := DB.QueryRow("SELECT password FROM stdusers where username = ?", username).Scan(&databasePassword)

	if err != nil {
		// todo: add a info box including "Username not existed"
		//http.Redirect(res, req, "/stdLogin", 301)
		http.Error(res,"Cannot Login",500)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword),
		[]byte(password))
	if err != nil {
		// todo: add a info box including "Wrong password"
		//http.Redirect(res, req, "/stdLogin", 301)
		http.Error(res,"Cannot Login",500)
		return
	}

	updateDatehandler, err := DB.Prepare("UPDATE stdusers SET lastlogindate = ? WHERE username = ?")
	if err != nil {
		// todo: add a info box to inform loginDate error
		//http.Redirect(res, req, "/stdLogin", 301)
		http.Error(res,"Cannot Login",500)
		return
	}

	logindate := time.Now().Local()
	_, err = updateDatehandler.Exec(logindate, username)
	if err != nil {
		//http.Redirect(res, req, "/stdLogin", 301)
		http.Error(res,"Cannot Login",500)
		return
	}

	//expiration := time.Now()
	//expiration = expiration.AddDate(0, 0, 7)
	//cookie := http.Cookie{Name: "username", Value: username, Expires: expiration}
	//fmt.Println("Setting cookies... cookie: ", cookie)

	//http.SetCookie(res, &cookie)

	session := Session{Username: username, Connection: true, SessionType:Student}
	(&session).StartSession()
	go (&session).EndSession()

	cookies := session.SetCookies()
	http.SetCookie(res, &cookies)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
